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
	errMsg := ""

	var number strings.Builder
	var literalNumber strings.Builder
	number.WriteByte(current)
	literalNumber.WriteByte(current)

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

		literalNumber.WriteByte(next)

		switch next {
		case '_':
			if lastByte == '_' {
				errMsg = "multiple underscores in number %s"
			}

			_, err = t.GetNext()

			if err != nil {
				errMsg = "invalid number %s"
			}

		case '.':
			if (numberFlags & NumberFlagFloat) != 0 {
				errMsg = "multiple decimal points in number %s"
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
				errMsg = "multiple exponent signs in number %s"
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
				break GETNEXT
			}

			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			if lastByte == '+' || lastByte == '-' {
				errMsg = "invalid number %s"
			}

			isNumberValid = false

			if next == '-' {
				number.WriteByte(next)
			}

		default:
			break GETNEXT
		}

		lastByte = next
	}

	if !isLastByteValid(lastByte) {
		errMsg = "invalid number %s"
	}

	if errMsg != "" {
		return token.Token{}, fmt.Errorf(errMsg, literalNumber.String())
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
