package parser

import "fmt"

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() error {
	tokens, err := p.tokenizer.Tokenize()

	if err != nil {
		return err
	}

	p.tokens = tokens
	p.tokenLen = len(tokens)

	nextToken, err := p.GetNextToken()

	if err != nil {
		return err
	}

	ast, err := p.parseExpr(nextToken, nil, 0)

	if err != nil {
		return err
	}

	fmt.Println(ast.Expr())

	return nil
}
