package tokenizer

import (
	"unicode"

	"github.com/Dobefu/pratt-parser/internal/errorutil"
	"github.com/Dobefu/pratt-parser/internal/token"
)

// Tokenize analyzes the expression string and turns it into tokens.
func (t *Tokenizer) Tokenize() ([]*token.Token, error) {
	approxNumTokens := (t.expLen / 3)
	tokens := make([]*token.Token, 0, approxNumTokens)

	for !t.isEOF {
		next, err := t.GetNext()

		if err != nil {
			return tokens, err
		}

		switch next {
		// Whitespace characters.
		case ' ', '\r', '\t':
			continue

		// Numeric characters.
		case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newToken, err := t.parseNumber(next)

			if err != nil {
				return nil, err
			}

			tokens = append(tokens, newToken)

		case '+':
			tokens = append(tokens, token.NewToken(
				"+",
				token.TokenTypeOperationAdd,
			))

		case '-':
			tokens = append(tokens, token.NewToken(
				"-",
				token.TokenTypeOperationSub,
			))

		case '*':
			token, err := t.handleAsterisk()

			if err != nil {
				return nil, err
			}

			tokens = append(tokens, token)

		case '%':
			tokens = append(tokens, token.NewToken(
				"%",
				token.TokenTypeOperationMod,
			))

		case '/':
			tokens = append(tokens, token.NewToken(
				"/",
				token.TokenTypeOperationDiv,
			))

		case '(':
			tokens = append(tokens, token.NewToken(
				"(",
				token.TokenTypeLParen,
			))

		case ')':
			tokens = append(tokens, token.NewToken(
				")",
				token.TokenTypeRParen,
			))

		case ',':
			tokens = append(tokens, token.NewToken(
				",",
				token.TokenTypeComma,
			))

		default:
			newToken, err := t.parseUnknownChar(next)

			if err != nil {
				return nil, err
			}

			tokens = append(tokens, newToken)
		}
	}

	return tokens, nil
}

func (t *Tokenizer) handleAsterisk() (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '*' {
		_, err = t.GetNext()

		if err != nil {
			return nil, err
		}

		return token.NewToken("**", token.TokenTypeOperationPow), nil
	}

	return token.NewToken("*", token.TokenTypeOperationMul), nil
}

func (t *Tokenizer) parseIdentifier(firstChar rune) (*token.Token, error) {
	identifier := string(firstChar)

	for !t.isEOF {
		next, err := t.Peek()

		if err != nil {
			break
		}

		if unicode.IsLetter(rune(next)) ||
			next == '_' ||
			unicode.IsDigit(rune(next)) {
			_, err = t.GetNext()

			if err != nil {
				return nil, err
			}

			identifier += string(next)

			continue
		}

		break
	}

	return token.NewToken(identifier, token.TokenTypeIdentifier), nil
}

func (t *Tokenizer) parseUnknownChar(next rune) (*token.Token, error) {
	if unicode.IsLetter(rune(next)) || next == '_' {
		return t.parseIdentifier(rune(next))
	}

	return nil, errorutil.NewError(
		errorutil.ErrorMsgUnexpectedChar,
		string(next),
		t.expIdx,
	)
}
