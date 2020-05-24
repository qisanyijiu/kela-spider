package parser

import (
	"github.com/PuerkitoBio/goquery"
	"kela-spider/types"
	"time"
)

type AlbumParser struct {
	Html  string
	ID    uint32
	doc   *goquery.Document
	name  string
	date  time.Time
	model types.Model
	tag   []*types.AlbumTag
}

var _ Parser = (*AlbumParser)(nil)

func NewAlbumParser(id uint32, html string) *AlbumParser {
	return &AlbumParser{
		Html: html,
		ID:   id,
	}
}

func (p *AlbumParser) Name() string {

}

func (p *AlbumParser) Date() time.Time {

}

func (p *AlbumParser) Model() types.Model {

}

func (p *AlbumParser) TitlePage() string {

}

func (p *AlbumParser) Tags() []*types.AlbumTag {

}

func (p *AlbumParser) Parse() (interface{}, error) {
}
