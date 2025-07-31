package parser

import (
	"log/slog"

	"github.com/Dobefu/pratt-parser/internal/token"
)

// Tokenize analyzes the expression string and turns it into tokens.
func (p *Parser) Tokenize() ([]token.Token, error) {
	var tokens []token.Token

	for !p.IsEOF() {
		next, err := p.GetNext()

		if err != nil {
			return tokens, err
		}

		slog.Info(string(next))
	}

	return tokens, nil
}
