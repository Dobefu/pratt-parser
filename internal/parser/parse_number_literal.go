package parser

import (
	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func (p *Parser) parseNumberLiteral(currentToken *token.Token) (ast.ExprNode, error) {
	return &ast.NumberLiteral{
		Value: currentToken.Atom,
	}, nil
}
