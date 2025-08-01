package parser

import (
	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func (p *Parser) parsePrefixExpr(currentToken *token.Token) (ast.ExprNode, error) {
	currentNumberLiteral, err := p.parseNumberLiteral(currentToken)

	if err != nil {
		return nil, err
	}

	return &ast.PrefixExpr{
		Operator: *currentToken,
		Operand:  currentNumberLiteral,
	}, nil
}
