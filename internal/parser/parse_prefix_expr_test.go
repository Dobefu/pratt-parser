package parser

import (
	"testing"
)

func TestParsePrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1",
			expected: "1",
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		ast, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got %v", err)

			continue
		}

		if ast.Expr() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, ast.Expr())
		}
	}
}

func TestParsePrefixExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
	}{
		{
			input: "+",
		},
		{
			input: "++",
		},
		{
			input: "(",
		},
		{
			input: ")",
		},
		{
			input: "(1 + 1",
		},
		{
			input: "(1 + 1 +",
		},
		{
			input: "(1 + 1 (",
		},
	}

	for _, test := range tests {
		_, err := NewParser(test.input).Parse()

		if err == nil {
			t.Errorf("expected error for %s, got none", test.input)
		}
	}
}
