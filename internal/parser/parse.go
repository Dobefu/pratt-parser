package parser

import (
	"github.com/Dobefu/pratt-parser/internal/ast"
)

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() error {
	tokens, err := p.tokenizer.Tokenize()

	if err != nil {
		return err
	}

	p.tokens = tokens
	p.tokenLen = len(tokens)

	var leftExpr ast.ExprNode

	for !p.isEOF {
		nextToken, err := p.GetNextToken()

		if err != nil {
			return err
		}

		expr, err := p.parseExpr(nextToken, leftExpr, 0)

		if err != nil {
			return err
		}

		leftExpr = expr
	}

	return nil
}
