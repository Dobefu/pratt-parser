package parser

import (
	"errors"
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

const maxRecursionDepth = 1000

func (p *Parser) parseExpr(
	currentToken *token.Token,
	leftExpr ast.ExprNode,
	recursionDepth int,
) (ast.ExprNode, error) {
	if recursionDepth > maxRecursionDepth {
		return nil, errors.New("maximum recursion depth reached")
	}

	switch currentToken.TokenType {
	case
		token.TokenTypeNumber:
		expr, err := p.parseNumberLiteral(currentToken)

		if err != nil {
			return nil, err
		}

		return expr, err

	case
		token.TokenTypeOperationAdd,
		token.TokenTypeOperationSub,
		token.TokenTypeOperationMul,
		token.TokenTypeOperationDiv:
		nextToken, err := p.PeekNextToken()

		if err != nil {
			return nil, err
		}

		if p.getBindingPower(currentToken) < p.getBindingPower(nextToken) {
			return p.parseExpr(nextToken, leftExpr, recursionDepth+1)
		}

		return p.parseBinaryExpr(currentToken, leftExpr, recursionDepth)

	case
		token.TokenTypeLParen:
		return p.parseExpr(currentToken, leftExpr, recursionDepth+1)

	default:
		return nil, fmt.Errorf("unexpected token: %s", currentToken.Atom)
	}
}
