package parser

type Parser interface {
	Parse() (interface{}, error)
}


