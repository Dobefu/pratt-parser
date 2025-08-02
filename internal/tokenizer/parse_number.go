package tokenizer

import (
	"fmt"
	"strings"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func (t *Tokenizer) parseNumber(current byte) (*token.Token, error) {
	var number strings.Builder
	number.WriteByte(current)

	lastByte := current
	isNumberValid := true
	isFloat := current == '.'

GETNEXT:
	for !t.isEOF {
		next, err := t.GetNext()

		if err != nil {
			return nil, err
		}

		switch next {
		case '_':
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
			t.expIdx--

			if t.expIdx < t.expLen {
				t.isEOF = false
			}

			break GETNEXT
		}

		lastByte = next
	}

	if lastByte == '.' || !isNumberValid {
		return nil, fmt.Errorf("invalid number %s", number.String())
	}

	return &token.Token{
		Atom:      number.String(),
		TokenType: token.TokenTypeNumber,
	}, nil
}
