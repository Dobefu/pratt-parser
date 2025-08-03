package parser

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func (p *Parser) parsePrefixExpr(
	currentToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	switch currentToken.TokenType {
	case token.TokenTypeNumber:
		return p.parseNumberLiteral(currentToken)

	case
		token.TokenTypeOperationAdd,
		token.TokenTypeOperationSub:
		nextToken, err := p.GetNextToken()
		if err != nil {
			return nil, err
		}

		operand, err := p.parseExpr(
			nextToken,
			nil,
			p.getBindingPower(currentToken, true),
			recursionDepth+1,
		)

		if err != nil {
			return nil, err
		}

		return &ast.PrefixExpr{
			Operator: *currentToken,
			Operand:  operand,
		}, nil

	case token.TokenTypeLParen:
		nextToken, err := p.GetNextToken()

		if err != nil {
			return nil, err
		}

		expr, err := p.parseExpr(nextToken, nil, 0, recursionDepth+1)

		if err != nil {
			return nil, err
		}

		rparenToken, err := p.GetNextToken()

		if err != nil {
			return nil, fmt.Errorf("expected ')', but got EOF")
		}

		if rparenToken.TokenType != token.TokenTypeRParen {
			return nil, fmt.Errorf("expected ')', got %s", rparenToken.Atom)
		}

		return expr, nil

	default:
		return nil, fmt.Errorf("unexpected token: %s", currentToken.Atom)
	}
}
