package parser

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func (p *Parser) parseFunctionCall(functionName string, recursionDepth int) (ast.ExprNode, error) {
	lparenToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if lparenToken.TokenType != token.TokenTypeLParen {
		return nil, fmt.Errorf("expected '(', got: %s", lparenToken.Atom)
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	var args []ast.ExprNode

	if nextToken.TokenType != token.TokenTypeRParen {
		args, err = p.parseFunctionCallArguments(recursionDepth + 1)

		if err != nil {
			return nil, err
		}
	}

	rparenToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if rparenToken.TokenType != token.TokenTypeRParen {
		return nil, fmt.Errorf("expected ')', got: %s", rparenToken.Atom)
	}

	return &ast.FunctionCall{
		FunctionName: functionName,
		Arguments:    args,
	}, nil
}

func (p *Parser) parseFunctionCallArguments(
	recursionDepth int,
) ([]ast.ExprNode, error) {
	// Pre-allocate the size of the slice, to reduce allocs.
	args := make([]ast.ExprNode, 0, 4)

	for {
		argToken, err := p.GetNextToken()

		if err != nil {
			return nil, err
		}

		arg, err := p.parseExpr(argToken, nil, 0, recursionDepth+1)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		nextToken, err := p.PeekNextToken()

		if err != nil {
			return nil, err
		}

		if nextToken.TokenType == token.TokenTypeRParen {
			break
		}

		if nextToken.TokenType != token.TokenTypeComma {
			return nil, fmt.Errorf("unexpected token: %s", nextToken.Atom)
		}

		_, err = p.GetNextToken()

		if err != nil {
			return nil, err
		}
	}

	return args, nil
}
