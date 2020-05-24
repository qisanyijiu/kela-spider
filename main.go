package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"sync"
	"time"
)

const (
	ALBUM_URL   = "https://www.kelagirls.com/albums/album-%d.html"
	ALBUM_START = 197
	ALBUM_END   = 198
)

func albums() {
	var (
		wg sync.WaitGroup
		ch = make(chan string)
		c  = http.Client{
			Timeout: 600 * time.Second,
		}
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go exec(ch, &c, &wg)
	}
	for i := ALBUM_START; i < ALBUM_END; i++ {
		url := fmt.Sprintf(ALBUM_URL, i)
		ch <- url
	}
	close(ch)
	wg.Wait()
}

func exec(urls chan string, httpClient *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		resp, err := httpClient.Get(url)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			fmt.Printf("URL: %40s, Http-Code: %4d", url, resp.StatusCode)
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		fmt.Println(doc)
	}
}

func main() {
	albums()
}
