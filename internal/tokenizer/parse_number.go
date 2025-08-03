package tokenizer

import (
	"fmt"
	"strings"

	"github.com/Dobefu/pratt-parser/internal/token"
)

// NumberFlags represents the type of number being parsed.
type NumberFlags byte

const (
	// NumberFlagFloat represents a floating point number.
	NumberFlagFloat NumberFlags = 1 << iota
	// NumberFlagExponent represents an exponent number.
	NumberFlagExponent
)

func (t *Tokenizer) parseNumber(current byte) (token.Token, error) {
	var number strings.Builder
	number.WriteByte(current)

	lastByte := current
	isNumberValid := true

	var numberFlags NumberFlags

	if current == '.' {
		numberFlags |= NumberFlagFloat
	}

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
			if (numberFlags & NumberFlagFloat) != 0 {
				isNumberValid = false
			}

			if lastByte == '.' {
				return token.Token{}, fmt.Errorf("invalid number %s", number.String())
			}

			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			number.WriteByte(next)
			numberFlags |= NumberFlagFloat

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if !isNumberValid && (numberFlags&NumberFlagExponent) != 0 {
				isNumberValid = true
			}

			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			number.WriteByte(next)

		case 'e', 'E':
			if !isNumberValid || (numberFlags&NumberFlagExponent) != 0 {
				return token.Token{}, fmt.Errorf("invalid number %s", number.String())
			}

			numberFlags |= NumberFlagExponent
			isNumberValid = false

			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			number.WriteByte(next)

		case '+', '-':
			if (numberFlags & NumberFlagExponent) == 0 {
				return token.Token{}, fmt.Errorf("invalid number %s", number.String())
			}

			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			if lastByte == '+' || lastByte == '-' {
				return token.Token{}, fmt.Errorf("invalid number %s", number.String()+"_")
			}

			isNumberValid = false

			if next != '+' {
				number.WriteByte(next)
			}

		default:
			break GETNEXT
		}

		lastByte = next
	}

	if !isLastByteValid(lastByte) {
		return token.Token{}, fmt.Errorf("invalid number %s", number.String())
	}

	return token.Token{
		Atom:      number.String(),
		TokenType: token.TokenTypeNumber,
	}, nil
}

func isLastByteValid(lastByte byte) bool {
	switch lastByte {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true

	default:
		return false
	}
}
