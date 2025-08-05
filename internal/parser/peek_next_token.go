package parser

import (
	"errors"

	"github.com/Dobefu/pratt-parser/internal/token"
)

// PeekNextToken gets the next token without advancing the current token index.
func (p *Parser) PeekNextToken() (*token.Token, error) {
	if p.isEOF {
		return nil, errors.New("unexpected end of expression")
	}

	return &p.tokens[p.tokenIdx], nil
}
