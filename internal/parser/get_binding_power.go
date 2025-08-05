package parser

import (
	"github.com/Dobefu/pratt-parser/internal/token"
)

const (
	bindingPowerParentheses    = 1000
	bindingPowerUnary          = 300
	bindingPowerPower          = 400
	bindingPowerMultiplicative = 200
	bindingPowerAdditive       = 100
	bindingPowerDefault        = 0
)

// getBindingPower returns the binding power of the current token.
func (p *Parser) getBindingPower(currentToken *token.Token, isUnary bool) int {
	switch currentToken.TokenType {
	case token.TokenTypeLParen, token.TokenTypeRParen:
		return bindingPowerParentheses

	case token.TokenTypeOperationPow, token.TokenTypeOperationMod:
		return bindingPowerPower

	case token.TokenTypeOperationMul, token.TokenTypeOperationDiv:
		return bindingPowerMultiplicative

	case token.TokenTypeOperationAdd, token.TokenTypeOperationSub:
		if isUnary {
			return bindingPowerUnary
		}

		return bindingPowerAdditive

	default:
		return bindingPowerDefault
	}
}
