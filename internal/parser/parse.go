package parser

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
)

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() (ast.ExprNode, error) {
	tokens, err := p.tokenizer.Tokenize()

	if err != nil {
		return nil, err
	}

	p.tokens = tokens
	p.tokenLen = len(tokens)
	p.isEOF = len(tokens) <= 0

	if len(tokens) == 0 {
		return nil, fmt.Errorf("no tokens to parse")
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	ast, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return nil, err
	}

	p.ast = ast

	return ast, nil
}
