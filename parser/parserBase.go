package parser

import "github.com/PuerkitoBio/goquery"

type ParserBase struct {
	Html string
	doc  *goquery.Document
}
