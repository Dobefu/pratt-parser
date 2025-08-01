// Package parser defines the actual Pratt parser.
package parser

import (
	"github.com/Dobefu/pratt-parser/internal/token"
	"github.com/Dobefu/pratt-parser/internal/tokenizer"
)

// Parser defines the parser itself.
type Parser struct {
	tokenizer *tokenizer.Tokenizer
	tokens    []token.Token
	tokenIdx  int
	tokenLen  int
	isEOF     bool
}

// NewParser creates a new instance of the Parser struct.
func NewParser(exp string) *Parser {
	return &Parser{
		tokenizer: tokenizer.NewTokenizer(exp),
		tokens:    []token.Token{},
		tokenIdx:  0,
		tokenLen:  0,
		isEOF:     false,
	}
}
