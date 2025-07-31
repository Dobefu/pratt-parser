package parser

import (
	"fmt"
	"strings"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func (p *Parser) parseNumber(current byte) (*token.Token, error) {
	var number strings.Builder
	number.WriteByte(current)

	lastByte := current
	isNumberValid := true
	isFloat := current == '.'

GETNEXT:
	for !p.IsEOF() {
		next, err := p.Peek()

		if err != nil {
			return nil, err
		}

		switch next {
		case '_':
			p.expIdx++

			continue GETNEXT

		case '.':
			if isFloat {
				isNumberValid = false
			}

			number.WriteByte(next)
			isFloat = true

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number.WriteByte(next)

		default:
			break GETNEXT
		}

		lastByte = next

		p.expIdx++
	}

	if lastByte == '.' || !isNumberValid {
		return nil, fmt.Errorf("invalid number %s", number.String())
	}

	return &token.Token{
		Atom:      number.String(),
		TokenType: token.TokenTypeNumber,
	}, nil
}
