package parser

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

const maxRecursionDepth = 100_000

func (p *Parser) parseExpr(
	currentToken *token.Token,
	leftExpr ast.ExprNode,
	minPrecedence int,
	recursionDepth int,
) (ast.ExprNode, error) {
	if recursionDepth > maxRecursionDepth {
		return nil, fmt.Errorf("maximum recursion depth of (%d) exceeded", maxRecursionDepth)
	}

	if leftExpr == nil {
		var err error

		leftExpr, err = p.parsePrefixExpr(currentToken, recursionDepth+1)

		if err != nil {
			return nil, err
		}
	}

	if p.isEOF {
		return leftExpr, nil
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	switch nextToken.TokenType {
	case
		token.TokenTypeOperationAdd,
		token.TokenTypeOperationSub,
		token.TokenTypeOperationMul,
		token.TokenTypeOperationDiv,
		token.TokenTypeOperationMod:

		return p.handleBasicOperatorTokens(
			nextToken,
			leftExpr,
			minPrecedence,
			recursionDepth,
		)

	case token.TokenTypeOperationPow:
		return p.handlePowToken(leftExpr, minPrecedence, recursionDepth)

	default:
		return leftExpr, nil
	}
}

func (p *Parser) handleBasicOperatorTokens(
	nextToken *token.Token,
	leftExpr ast.ExprNode,
	minPrecedence int,
	recursionDepth int,
) (ast.ExprNode, error) {
	if p.getBindingPower(nextToken, false) < minPrecedence {
		return leftExpr, nil
	}

	operator, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	rightToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	expr, err := p.parseBinaryExpr(operator, leftExpr, rightToken, recursionDepth+1)

	if err != nil {
		return nil, err
	}

	return p.parseExpr(nil, expr, minPrecedence, recursionDepth+1)
}

func (p *Parser) handlePowToken(
	leftExpr ast.ExprNode,
	minPrecedence int,
	recursionDepth int,
) (ast.ExprNode, error) {
	operator, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	rightToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	expr, err := p.parseBinaryExpr(operator, leftExpr, rightToken, recursionDepth+1)

	if err != nil {
		return nil, err
	}

	return p.parseExpr(nil, expr, minPrecedence-1, recursionDepth+1)
}
