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
			errMsg = t.handleUnderscore(lastByte, errMsg)

		case '.':
			errMsg = t.handleDecimalPoint(&numberFlags, &number, errMsg)

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			errMsg = t.handleDigit(
				&numberFlags,
				&number,
				&isNumberValid,
				next,
				errMsg,
			)

		case 'e', 'E':
			errMsg = t.handleExponent(
				&numberFlags,
				&number,
				&isNumberValid,
				next,
				errMsg,
			)

		case '+', '-':
			var shouldBreak bool

			errMsg, shouldBreak = t.handleAdditionAndSubtraction(
				&numberFlags,
				&number,
				&isNumberValid,
				lastByte,
				next,
				errMsg,
			)

			if shouldBreak {
				break GETNEXT
			}

		default:
			break GETNEXT
		}

		_, err = t.GetNext()

		if err != nil {
			return token.Token{}, err
		}

		lastByte = next
	}

	if !isLastByteValid(lastByte) {
		errMsg = "trailing character in number %s"
	}

	if errMsg != "" {
		return token.Token{}, fmt.Errorf(errMsg, literalNumber.String())
	}

	return token.Token{
		Atom:      number.String(),
		TokenType: token.TokenTypeNumber,
	}, nil
}

func (t *Tokenizer) handleUnderscore(
	lastByte byte,
	currentErrMsg string,
) string {
	if lastByte == '_' {
		return "multiple underscores in number %s"
	}

	return currentErrMsg
}

func (t *Tokenizer) handleDecimalPoint(
	numberFlags *NumberFlags,
	number *strings.Builder,
	currentErrMsg string,
) string {
	if (*numberFlags & NumberFlagFloat) != 0 {
		return "multiple decimal points in number %s"
	}

	*numberFlags |= NumberFlagFloat
	number.WriteByte('.')

	return currentErrMsg
}

func (t *Tokenizer) handleDigit(
	numberFlags *NumberFlags,
	number *strings.Builder,
	isNumberValid *bool,
	next byte,
	currentErrMsg string,
) string {
	if !*isNumberValid && (*numberFlags&NumberFlagExponent) != 0 {
		*isNumberValid = true
	}

	number.WriteByte(next)

	return currentErrMsg
}

func (t *Tokenizer) handleExponent(
	numberFlags *NumberFlags,
	number *strings.Builder,
	isNumberValid *bool,
	next byte,
	currentErrMsg string,
) string {
	if !*isNumberValid || (*numberFlags&NumberFlagExponent) != 0 {
		return "multiple exponent signs in number %s"
	}

	*numberFlags |= NumberFlagExponent
	*isNumberValid = false

	number.WriteByte(next)

	return currentErrMsg
}

func (t *Tokenizer) handleAdditionAndSubtraction(
	numberFlags *NumberFlags,
	number *strings.Builder,
	isNumberValid *bool,
	lastByte byte,
	next byte,
	currentErrMsg string,
) (string, bool) {
	if (*numberFlags & NumberFlagExponent) == 0 {
		return currentErrMsg, true
	}

	if lastByte == '+' || lastByte == '-' {
		return "multiple addition or subtraction signs in exponent %s", true
	}

	*isNumberValid = false

	if next == '-' {
		number.WriteByte(next)
	}

	return currentErrMsg, false
}

func isLastByteValid(lastByte byte) bool {
	switch lastByte {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true

	default:
		return false
	}
}
