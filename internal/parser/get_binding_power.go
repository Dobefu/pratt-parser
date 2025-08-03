package parser

import (
	"github.com/Dobefu/pratt-parser/internal/token"
)

// getBindingPower returns the binding power of the current token.
func (p *Parser) getBindingPower(currentToken *token.Token, isUnary bool) int {
	switch currentToken.TokenType {
	case token.TokenTypeLParen, token.TokenTypeRParen:
		return 1000

	case token.TokenTypeOperationPow, token.TokenTypeOperationMod:
		return 400

	case token.TokenTypeOperationMul, token.TokenTypeOperationDiv:
		return 200

	case token.TokenTypeOperationAdd, token.TokenTypeOperationSub:

		if isUnary {
			return 300
		}

		return 100

	default:
		return 0
	}
}
