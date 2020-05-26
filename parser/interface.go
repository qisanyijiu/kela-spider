package parser

type Parser interface {
	Parse(res interface{}) error
}
