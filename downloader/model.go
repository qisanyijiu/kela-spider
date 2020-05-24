package downloader

import (
	"fmt"
	"io/ioutil"
	"kela-spider/parser"
	"net/http"
	"sync"
	"time"
)

const (
	MD_TPL  = ""
	MD_B_ID = 1
	MD_E_ID = 200
)

type ModelDownloader struct {
	tableName string
	urlTpl    string
	beginId   int
	endId     int
	state     State
	err       []error
	hc        *http.Client
}

func NewModelDownloader(url string, beginId, endId int, table string) *ModelDownloader {
	return &ModelDownloader{
		tableName: table,
		urlTpl:    url,
		beginId:   beginId,
		endId:     endId,
		state:     waiting,
		err:       []error{},
		hc: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       600 * time.Second,
		},
	}
}

func (m *ModelDownloader) Download() {
	var (
		err       error
		ids       = make(chan int)
		parserChs = []chan *parser.ModelParser{}
	)
	for i := 0; i < 10; i++ {
		parserChs = append(parserChs, m.Exec(ids))
	}
	go func() {
		for i := m.beginId; i <= m.endId; i ++ {
			ids <- i
		}
	}()
	parserCh := modelFanIn(parserChs...)
	for p := range parserCh {
		_, err = p.Parse()
		if err != nil {
			m.err = append(m.err, err)
		}
	}
}

func (m *ModelDownloader) GetState() string {
	return m.state.Desc()
}

func (m *ModelDownloader) GetError() []error {
	return m.err
}

func (m *ModelDownloader) Exec(ch chan int) chan *parser.ModelParser {
	var parserCh = make(chan *parser.ModelParser)
	go func() {
		for id := range ch {
			url := fmt.Sprintf(m.urlTpl, id)
			html, err := download(url, m.hc)
			if err != nil {
				parserCh <- parser.NewModelParser(uint32(id), html)
			}
		}
		close(ch)
	}()
	return parserCh
}

func modelFanIn(chs ...chan *parser.ModelParser) chan *parser.ModelParser {
	var (
		wg     sync.WaitGroup
		result = make(chan *parser.ModelParser)
	)
	output := func(ch <-chan *parser.ModelParser) {
		for n := range ch {
			result <- n
		}
		wg.Done()
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go output(ch)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

func download(url string, httpClient *http.Client) (string, error) {
	var err error = nil
	for i := 0; i < 10; i++ {
		resp, err := httpClient.Get(url)
		defer resp.Body.Close()
		if err != nil && resp != nil {
			switch resp.StatusCode {
			case http.StatusOK:
				b, _ := ioutil.ReadAll(resp.Body)
				return string(b), nil
			default:
				err = fmt.Errorf("http code err: %d", resp.StatusCode)
				time.Sleep(10 * time.Second)
				continue
			}
		}
		time.Sleep(10 * time.Second)
	}
	return "", err
}
