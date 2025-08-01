package parser

import (
	"errors"

	"github.com/Dobefu/pratt-parser/internal/token"
)

// GetNextToken gets the next token and advances the current token index.
func (p *Parser) GetNextToken() (*token.Token, error) {
	if p.isEOF {
		return nil, errors.New("cannot get next token after EOF")
	}

	next := p.tokens[p.tokenIdx]
	p.tokenIdx++

	if p.tokenIdx >= p.tokenLen {
		p.isEOF = true
	}

	return &next, nil
}
