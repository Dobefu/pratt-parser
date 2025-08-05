// Package parser defines the actual Pratt parser.
package parser

import (
	"github.com/Dobefu/pratt-parser/internal/token"
)

// Parser defines the parser itself.
type Parser struct {
	tokens   []token.Token
	tokenIdx int
	tokenLen int
	isEOF    bool
}

// NewParser creates a new instance of the Parser struct.
func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens:   tokens,
		tokenIdx: 0,
		tokenLen: len(tokens),
		isEOF:    len(tokens) <= 0,
	}
}
