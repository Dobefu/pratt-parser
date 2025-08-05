package parser

import (
	"github.com/Dobefu/pratt-parser/internal/errorutil"
	"github.com/Dobefu/pratt-parser/internal/token"
)

// PeekNextToken gets the next token without advancing the current token index.
func (p *Parser) PeekNextToken() (*token.Token, error) {
	if p.isEOF {
		return nil, errorutil.NewError(errorutil.ErrorMsgUnexpectedEOF)
	}

	return &p.tokens[p.tokenIdx], nil
}
