package parser

import "github.com/PuerkitoBio/goquery"

type ModelParser struct {
	Html string
	ID   uint32
	doc  *goquery.Document
}

var _ Parser = (*ModelParser)(nil)

func NewModelParser(id uint32, html string) *ModelParser {
	return &ModelParser{
		Html: html,
		ID:   id,
	}
}

func (ModelParser) Name() string {

}

func (p *ModelParser) Height() uint8 {

}

func (p *ModelParser) Weight() uint8 {

}

func (p *ModelParser) Cup() string {

}

func (p *ModelParser) Bust() uint8 {

}

func (p *ModelParser) Waistline() uint8 {

}

func (p *ModelParser) Hips() uint8 {

}

func (p *ModelParser) City() string {

}

func (p *ModelParser) Parse() (interface{}, error) {

}
