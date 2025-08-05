package tokenizer

import (
	"github.com/Dobefu/pratt-parser/internal/charutil"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
	"github.com/Dobefu/pratt-parser/internal/token"
)

// Tokenize analyzes the expression string and turns it into tokens.
func (t *Tokenizer) Tokenize() ([]token.Token, error) {
	approxNumTokens := (t.expLen / 3)
	tokens := make([]token.Token, 0, approxNumTokens)

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
			tokens = append(tokens, token.Token{
				Atom:      "+",
				TokenType: token.TokenTypeOperationAdd,
			})

		case '-':
			tokens = append(tokens, token.Token{
				Atom:      "-",
				TokenType: token.TokenTypeOperationSub,
			})

		case '*':
			token, err := t.handleAsterisk()

			if err != nil {
				return nil, err
			}

			tokens = append(tokens, token)

		case '%':
			tokens = append(tokens, token.Token{
				Atom:      "%",
				TokenType: token.TokenTypeOperationMod,
			})

		case '/':
			tokens = append(tokens, token.Token{
				Atom:      "/",
				TokenType: token.TokenTypeOperationDiv,
			})

		case '(':
			tokens = append(tokens, token.Token{
				Atom:      "(",
				TokenType: token.TokenTypeLParen,
			})

		case ')':
			tokens = append(tokens, token.Token{
				Atom:      ")",
				TokenType: token.TokenTypeRParen,
			})

		case ',':
			tokens = append(tokens, token.Token{
				Atom:      ",",
				TokenType: token.TokenTypeComma,
			})

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

func (t *Tokenizer) handleAsterisk() (token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return token.Token{}, err
	}

	if next == '*' {
		_, err = t.GetNext()

		if err != nil {
			return token.Token{}, err
		}

		return token.Token{
			Atom:      "**",
			TokenType: token.TokenTypeOperationPow,
		}, nil
	}

	return token.Token{
		Atom:      "*",
		TokenType: token.TokenTypeOperationMul,
	}, nil
}

func (t *Tokenizer) parseIdentifier(firstChar rune) (token.Token, error) {
	identifier := string(firstChar)

	for !t.isEOF {
		next, err := t.Peek()

		if err != nil {
			break
		}

		if charutil.IsLetter(rune(next)) ||
			next == '_' ||
			charutil.IsDigit(rune(next)) {
			_, err = t.GetNext()

			if err != nil {
				return token.Token{}, err
			}

			identifier += string(next)

			continue
		}

		break
	}

	return token.Token{
		Atom:      identifier,
		TokenType: token.TokenTypeIdentifier,
	}, nil
}

func (t *Tokenizer) parseUnknownChar(next rune) (token.Token, error) {
	if charutil.IsLetter(rune(next)) || next == '_' {
		return t.parseIdentifier(rune(next))
	}

	return token.Token{}, errorutil.NewError(
		errorutil.ErrorMsgUnexpectedChar,
		string(next),
		t.expIdx,
	)
}
