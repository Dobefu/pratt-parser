package tokenizer

import (
	"fmt"
	"strings"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func (t *Tokenizer) parseNumber(current byte) (token.Token, error) {
	var number strings.Builder
	number.WriteByte(current)

	lastByte := current
	isNumberValid := true
	isFloat := current == '.'

GETNEXT:
	for !t.isEOF {
		next, err := t.Peek()

		if err != nil {
			return token.Token{}, err
		}

		switch next {
		case '_':
			if lastByte == '_' {
				return token.Token{}, fmt.Errorf("invalid number %s", number.String()+"_")
			}

			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			lastByte = next

			continue GETNEXT

		case '.':
			if isFloat {
				isNumberValid = false
			}

			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			number.WriteByte(next)
			isFloat = true

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			number.WriteByte(next)

		default:
			break GETNEXT
		}

		lastByte = next
	}

	if lastByte == '.' || lastByte == '_' || !isNumberValid {
		return token.Token{}, fmt.Errorf("invalid number %s", number.String())
	}

	return token.Token{
		Atom:      number.String(),
		TokenType: token.TokenTypeNumber,
	}, nil
}
