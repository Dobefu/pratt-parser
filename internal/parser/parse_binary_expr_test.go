package parser

import (
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
	"github.com/Dobefu/pratt-parser/internal/tokenizer"
)

func TestParseBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected *ast.BinaryExpr
	}{
		{
			input: "1 + 1",
			expected: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "1"},
				Right: &ast.NumberLiteral{Value: "1"},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
			},
		},
	}

	for _, test := range tests {
		to := tokenizer.NewTokenizer(test.input)
		tokens, _ := to.Tokenize()

		parser := NewParser(tokens)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got \"%v\"", err)

			continue
		}

		if result.Expr() != test.expected.Expr() {
			t.Errorf("expected \"%s\", got \"%s\"", test.expected.Expr(), result.Expr())
		}
	}
}

func TestParseBinaryExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []token.Token
		expected string
	}{
		{
			input:    []token.Token{},
			expected: "no tokens to parse",
		},
		{
			input: []token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: "cannot get next token after EOF",
		},
		{
			input: []token.Token{
				{Atom: "3", TokenType: token.TokenTypeNumber},
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
			},
			expected: "cannot get next token after EOF",
		},
	}

	for _, test := range tests {
		_, err := NewParser(test.input).Parse()

		if err == nil {
			t.Fatalf("expected error for \"%v\", got none", test.input)
		}

		if err.Error() != test.expected {
			t.Errorf(
				"expected error \"%s\", got \"%s\"",
				test.expected,
				err.Error(),
			)
		}
	}
}

func BenchmarkParseBinaryExpr(b *testing.B) {
	for b.Loop() {
		p := NewParser([]token.Token{})

		_, _ = p.parseBinaryExpr(
			&token.Token{
				Atom:      "1",
				TokenType: token.TokenTypeOperationAdd,
			},
			nil,
			nil,
			0,
		)
	}
}
