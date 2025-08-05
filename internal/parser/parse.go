package parser

import (
	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
)

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() (ast.ExprNode, error) {
	if len(p.tokens) <= 0 {
		return nil, errorutil.NewError(errorutil.ErrorMsgEmptyExpression)
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	ast, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return nil, err
	}

	return ast, nil
}
