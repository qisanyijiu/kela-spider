package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type ModelParser struct {
	*ParserBase
	ID   uint32
}

var _ Parser = (*ModelParser)(nil)

func NewModelParser(id uint32, html string) *ModelParser {
	buf := bytes.NewBufferString(html)
	doc, _ := goquery.NewDocumentFromReader(buf)
	return &ModelParser{
		ParserBase: &ParserBase{
			Html: html,
			doc:  doc,
		},
		ID:         id,
	}
}

func (p *ModelParser) Parse(res interface{}) (error) {
	return nil
}

func (p *ModelParser) Name() string {
	node := p.doc.Find(".dataright").First().Children().First()
	if node == nil{
		return ""
	}
	name := strings.Replace(node.Text(), "\n", "", -1)
	name = strings.Replace(name, " ", "", -1)
	return name
}

func (p *ModelParser) Height() uint8 {
	node := p.doc.Find(".dataright").First().Children().First().Next().Next()
	tmp := strings.Split(node.Text(), ":")
	if len(tmp) != 2 {
		return 0
	}
	hStr := strings.Replace(tmp[1], "CM", "", -1)
	height, err := strconv.ParseInt(hStr, 10, 64)
	if err != nil{
		return 0
	}
	return uint8(height)
}

func (p *ModelParser) Weight() uint8 {
	node := p.doc.Find(".mymiaoshu").First().Prev().Prev()
	if node == nil{
		return 0
	}
	tmp := strings.Split(node.Text(), ":")
	if len(tmp) != 2 {
		return 0
	}
	wStr := strings.Replace(tmp[1], "KG", "", -1)
	weight, err := strconv.ParseInt(wStr, 10, 64)
	if err != nil{
		return 0
	}
	return uint8(weight)
}

func (p *ModelParser) Cup() string {
	node := p.doc.Find(".mymiaoshu").First().Prev().Prev().Prev()
	if node == nil{
		return ""
	}
	tmp := strings.Split(node.Text(), ":")
	if len(tmp) != 2 {
		return ""
	}
	return tmp[1]
}

func (p *ModelParser)BWH() (uint8, uint8, uint8) {
	var (
		breast uint8
		waist uint8
		hip uint8
		)
	node := p.doc.Find(".dataright").First().Children().First().Next()
	if node == nil{
		return breast,waist,hip
	}
	tmp := strings.Split(node.Text(), ":")
	if len(tmp) != 2 {
		return breast,waist,hip
	}
	bwh := strings.Split(tmp[1], "-")
	b, err := strconv.ParseInt(bwh[0], 10, 64)
	if err != nil {
		return breast,waist,hip
	}
	w, err := strconv.ParseInt(bwh[1], 10, 64)
	if err != nil {
		return breast,waist,hip
	}
	h, err := strconv.ParseInt(bwh[2], 10, 64)
	if err != nil {
		return breast,waist,hip
	}
	breast,waist,hip = uint8(b), uint8(w), uint8(h)
	return breast,waist,hip
}

func (p *ModelParser) City() string {
	node := p.doc.Find(".dataright").First().Children().First().Next().Next().Next()
	tmp := strings.Split(node.Text(), ":")
	if len(tmp) != 2 {
		return ""
	}
	return tmp[1]
}
