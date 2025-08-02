package parser

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func (p *Parser) parseBinaryExpr(
	operatorToken *token.Token,
	leftExpr ast.ExprNode,
	rightToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	if leftExpr == nil {
		return nil, fmt.Errorf("unexpected token %s", operatorToken.Atom)
	}

	rightExpr, err := p.parseExpr(
		rightToken,
		nil,
		p.getBindingPower(operatorToken, false)+1,
		recursionDepth+1,
	)

	if err != nil {
		return nil, err
	}

	return &ast.BinaryExpr{
		Left:     leftExpr,
		Right:    rightExpr,
		Operator: *operatorToken,
	}, nil
}
