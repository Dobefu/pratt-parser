package tokenizer

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/token"
)

// Tokenize analyzes the expression string and turns it into tokens.
func (t *Tokenizer) Tokenize() ([]token.Token, error) {
	approxNumTokens := (t.expLen / 2) + 1
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
			nextChar, err := t.Peek()

			if err != nil {
				return nil, err
			}

			if nextChar == '*' {
				_, err = t.GetNext()

				if err != nil {
					return nil, err
				}

				tokens = append(tokens, token.Token{
					Atom:      "**",
					TokenType: token.TokenTypeOperationPow,
				})
			} else {
				tokens = append(tokens, token.Token{
					Atom:      "*",
					TokenType: token.TokenTypeOperationMul,
				})
			}

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

		default:
			return tokens, fmt.Errorf(
				"unexpected character %s at index %d",
				string(next),
				t.expIdx,
			)
		}
	}

	return tokens, nil
}
