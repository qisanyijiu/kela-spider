package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"kela-spider/types"
)

type ModelListParser struct {
	*ParserBase
}

var _ Parser = (*ModelListParser)(nil)

func NewModelListParser(html string) *ModelListParser {
	buf := bytes.NewBufferString(html)
	doc, _ := goquery.NewDocumentFromReader(buf)
	return &ModelListParser{
		ParserBase: &ParserBase{
			Html: html,
			doc:  doc,
		},
	}
}

func (p *ModelListParser) Parse(res interface{}) error {
	var (
		err error
		result = []*types.Model{}
	)

	container := p.doc.Find(".models").First()
	node := container.Children().First()
	for  {
		if node == nil{
			break
		}
		modelId, ok := node.Attr("modelid")
		if !ok {
			node = node.Next()
			continue
		}
		imgNode := node.Find(".headimg").First()
		imgUri, ok := imgNode.Attr("src")
		fmt.Printf("ModelID: %s\n", modelId)
		fmt.Printf("HeadImage: %s\n", imgUri)
		modelInfo := node.Find(".down").First().Children()
		name := modelInfo.First()
		fmt.Printf("name: %s\n", name.Text())
		bwh := name.Next()
		fmt.Printf("bwh: %s\n", bwh.Text())
		height := bwh.Next()
		fmt.Printf("height: %s\n", height.Text())
		node = node.Next()
		break
	}
	res = result
	return err
}
