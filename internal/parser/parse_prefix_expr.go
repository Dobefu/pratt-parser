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

	case token.TokenTypeOperationAdd, token.TokenTypeOperationSub:
		return p.parseUnaryOperator(currentToken, recursionDepth)

	case token.TokenTypeLParen:
		return p.parseParenthesizedExpr(recursionDepth)

	case token.TokenTypeIdentifier:
		return p.parseIdentifier(currentToken, recursionDepth)

	default:
		return nil, fmt.Errorf("unexpected token: %s", currentToken.Atom)
	}
}

func (p *Parser) parseUnaryOperator(
	operatorToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	nextToken, err := p.GetNextToken()
	if err != nil {
		return nil, err
	}

	operand, err := p.parseExpr(
		nextToken,
		nil,
		p.getBindingPower(operatorToken, true),
		recursionDepth+1,
	)

	if err != nil {
		return nil, err
	}

	return &ast.PrefixExpr{
		Operator: *operatorToken,
		Operand:  operand,
	}, nil
}

func (p *Parser) parseParenthesizedExpr(
	recursionDepth int,
) (ast.ExprNode, error) {
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
		return nil, fmt.Errorf("expected ')', got: %s", rparenToken.Atom)
	}

	return expr, nil
}

func (p *Parser) parseIdentifier(
	identifierToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeLParen {
		return p.parseFunctionCall(identifierToken.Atom, recursionDepth+1)
	}

	return nil, fmt.Errorf("identifiers are not yet supported: %s", identifierToken.Atom)
}
