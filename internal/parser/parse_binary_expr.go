package parser

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func (p *Parser) parseBinaryExpr(
	operatorToken *token.Token,
	leftExpr ast.ExprNode,
	recursionDepth int,
) (ast.ExprNode, error) {
	if leftExpr == nil {
		return nil, fmt.Errorf("unexpected token %s", operatorToken.Atom)
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	nextNumberLiteral, err := p.parseExpr(nextToken, leftExpr, recursionDepth+1)

	if err != nil {
		return nil, err
	}

	expr := &ast.BinaryExpr{
		Left:     leftExpr,
		Right:    nextNumberLiteral,
		Operator: *operatorToken,
	}

	return expr, nil
}
