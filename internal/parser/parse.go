package parser

import (
	"fmt"
)

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() (float64, error) {
	tokens, err := p.tokenizer.Tokenize()

	if err != nil {
		return 0, err
	}

	p.tokens = tokens
	p.tokenLen = len(tokens)
	p.isEOF = len(tokens) <= 0

	if len(tokens) == 0 {
		return 0, fmt.Errorf("no tokens to parse")
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return 0, err
	}

	ast, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return 0, err
	}

	result, err := p.evaluator.Evaluate(ast)

	if err != nil {
		return 0, err
	}

	return result, nil
}
