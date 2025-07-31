package parser

import (
	"fmt"
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

		switch next {
		// Whitespace characters.
		case ' ', '\r', '\t':
			continue

		// Numeric characters.
		case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newToken, err := p.parseNumber(next)

			if err != nil {
				return nil, err
			}

			tokens = append(tokens, *newToken)

		default:
			return tokens, fmt.Errorf("unexpected character %s", string(next))
		}

		slog.Info(string(next))
	}

	return tokens, nil
}
