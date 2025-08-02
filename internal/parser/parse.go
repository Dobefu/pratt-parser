package parser

import (
	"fmt"
	"log/slog"
)

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() error {
	tokens, err := p.tokenizer.Tokenize()

	if err != nil {
		return err
	}

	p.tokens = tokens
	p.tokenLen = len(tokens)
	p.isEOF = len(tokens) <= 0

	if len(tokens) == 0 {
		return fmt.Errorf("no tokens to parse")
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return err
	}

	ast, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return err
	}

	p.ast = ast

	if ast != nil {
		slog.Info(ast.Expr())
	}

	return nil
}
