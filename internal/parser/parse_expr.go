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

	if p.isEOF {
		return leftExpr, nil
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	switch currentToken.TokenType {
	case
		token.TokenTypeNumber:
		expr, err := p.parseNumberLiteral(currentToken)

		if err != nil {
			return nil, err
		}

		_, err = p.GetNextToken()

		if err != nil {
			return nil, err
		}

		return p.parseExpr(nextToken, expr, recursionDepth+1)

	case
		token.TokenTypeOperationAdd,
		token.TokenTypeOperationSub:
		if leftExpr == nil {
			_, err = p.GetNextToken()

			if err != nil {
				return nil, err
			}

			return p.parsePrefixExpr(nextToken, recursionDepth)
		}

		fallthrough

	case
		token.TokenTypeOperationMul,
		token.TokenTypeOperationDiv:

		if p.getBindingPower(currentToken) < p.getBindingPower(nextToken) {
			_, err = p.GetNextToken()

			if err != nil {
				return nil, err
			}

			return p.parseExpr(nextToken, leftExpr, recursionDepth+1)
		}

		_, err = p.GetNextToken()

		if err != nil {
			return nil, err
		}

		return p.parseBinaryExpr(nextToken, leftExpr, recursionDepth)

	case
		token.TokenTypeLParen:
		_, err = p.GetNextToken()

		if err != nil {
			return nil, err
		}

		return p.parseExpr(nextToken, leftExpr, recursionDepth+1)

	default:
		return nil, fmt.Errorf("unexpected token: %s", currentToken.Atom)
	}
}
