package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"kela-spider/constans"
	"strings"
	"time"
)

type AlbumParser struct {
	*ParserBase
	ID uint32
}

var _ Parser = (*AlbumParser)(nil)

func NewAlbumParser(id uint32, html string) *AlbumParser {
	buf := bytes.NewBufferString(html)
	doc, _ := goquery.NewDocumentFromReader(buf)
	return &AlbumParser{
		ParserBase: &ParserBase{
			Html: html,
			doc:  doc,
		},
		ID: id,
	}
}

func (p *AlbumParser) Parse(res interface{}) error {
	return nil
}

func (p *AlbumParser) Name() string {
	node := p.doc.Find(".xlright").First().Children().First()
	if node == nil {
		return ""
	}
	name := strings.Replace(node.Text(), "《", "", -1)
	name = strings.Replace(name, "》", "", -1)
	return name
}

func (p *AlbumParser) Date() time.Time {
	node := p.doc.Find("kela-date").First()
	if node == nil {
		return time.Time{}
	}
	t, err := time.ParseInLocation(constans.DATE, node.Text(), time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}

func (p *AlbumParser) CoverPic() string {
	return ""
}
