package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

type ModelAlbumListParser struct {
	*ParserBase
}

var _ Parser = (*ModelListParser)(nil)

func NewModelAlbumListParser(html string) *ModelAlbumListParser {
	buf := bytes.NewBufferString(html)
	doc, _ := goquery.NewDocumentFromReader(buf)
	return &ModelAlbumListParser{
		ParserBase: &ParserBase{
			Html: html,
			doc:  doc,
		},
	}
}

func (p *ModelAlbumListParser) Parse(res interface{}) error {
	return nil
}
