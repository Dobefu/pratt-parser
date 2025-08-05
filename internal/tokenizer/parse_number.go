package tokenizer

import (
	"strings"

	"github.com/Dobefu/pratt-parser/internal/errorutil"
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

func (t *Tokenizer) parseNumber(current rune) (token.Token, error) {
	var errMsg errorutil.ErrorMsg

	var number strings.Builder
	var literalNumber strings.Builder
	number.WriteRune(current)
	literalNumber.WriteRune(current)

	lastChar := current
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

		literalNumber.WriteRune(next)

		switch next {
		case '_':
			errMsg = t.handleUnderscore(lastChar, errMsg)

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
				lastChar,
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

		lastChar = next
	}

	if !isLastCharValid(lastChar) {
		errMsg = errorutil.ErrorMsgNumberTrailingChar
	}

	if errMsg != "" {
		return token.Token{}, errorutil.NewError(errMsg, literalNumber.String())
	}

	return token.Token{
		Atom:      number.String(),
		TokenType: token.TokenTypeNumber,
	}, nil
}

func (t *Tokenizer) handleUnderscore(
	lastChar rune,
	currentErrMsg errorutil.ErrorMsg,
) errorutil.ErrorMsg {
	if lastChar == '_' {
		return errorutil.ErrorMsgNumberMultipleUnderscores
	}

	return currentErrMsg
}

func (t *Tokenizer) handleDecimalPoint(
	numberFlags *NumberFlags,
	number *strings.Builder,
	currentErrMsg errorutil.ErrorMsg,
) errorutil.ErrorMsg {
	if (*numberFlags & NumberFlagFloat) != 0 {
		return errorutil.ErrorMsgNumberMultipleDecimalPoints
	}

	*numberFlags |= NumberFlagFloat
	number.WriteRune('.')

	return currentErrMsg
}

func (t *Tokenizer) handleDigit(
	numberFlags *NumberFlags,
	number *strings.Builder,
	isNumberValid *bool,
	next rune,
	currentErrMsg errorutil.ErrorMsg,
) errorutil.ErrorMsg {
	if !*isNumberValid && (*numberFlags&NumberFlagExponent) != 0 {
		*isNumberValid = true
	}

	number.WriteRune(next)

	return currentErrMsg
}

func (t *Tokenizer) handleExponent(
	numberFlags *NumberFlags,
	number *strings.Builder,
	isNumberValid *bool,
	next rune,
	currentErrMsg errorutil.ErrorMsg,
) errorutil.ErrorMsg {
	if !*isNumberValid || (*numberFlags&NumberFlagExponent) != 0 {
		return errorutil.ErrorMsgNumberMultipleExponentSigns
	}

	*numberFlags |= NumberFlagExponent
	*isNumberValid = false

	number.WriteRune(next)

	return currentErrMsg
}

func (t *Tokenizer) handleAdditionAndSubtraction(
	numberFlags *NumberFlags,
	number *strings.Builder,
	isNumberValid *bool,
	lastChar rune,
	next rune,
	currentErrMsg errorutil.ErrorMsg,
) (errorutil.ErrorMsg, bool) {
	if (*numberFlags & NumberFlagExponent) == 0 {
		return currentErrMsg, true
	}

	if lastChar == '+' || lastChar == '-' {
		return errorutil.ErrorMsgNumberMultipleConsecutiveExponentSigns, true
	}

	*isNumberValid = false

	if next == '-' {
		number.WriteRune(next)
	}

	return currentErrMsg, false
}

func isLastCharValid(lastChar rune) bool {
	switch lastChar {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true

	default:
		return false
	}
}
