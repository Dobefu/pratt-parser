// Package parser defines the actual Pratt parser.
package parser

// Parser defines the parser itself.
type Parser struct {
	exp    string
	expLen int
	expIdx uint
	isEOF  bool
}

// NewParser creates a new instance of the Parser struct.
func NewParser(exp string) *Parser {
	return &Parser{
		exp:    exp,
		expLen: len(exp),
		expIdx: 0,
	}
}
