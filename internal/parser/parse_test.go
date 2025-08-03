package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1",
			expected: "1",
		},
		{
			input:    "1 + 2",
			expected: "(1 + 2)",
		},
		{
			input:    "1 ** 1",
			expected: "(1 ** 1)",
		},
		{
			input:    "1 % 1",
			expected: "(1 % 1)",
		},
		{
			input:    "1 + 2 * 3",
			expected: "(1 + (2 * 3))",
		},
		{
			input:    "1 + 2 * 3 / 4",
			expected: "(1 + ((2 * 3) / 4))",
		},
		{
			input:    "1 + 2 * 3 / 4 - 5",
			expected: "((1 + ((2 * 3) / 4)) - 5)",
		},
		{
			input:    "(1 + 10 / 5 + 2 * 5 - ( -123 - (-(5))))",
			expected: "(((1 + (10 / 5)) + (2 * 5)) - ((- 123) - (- 5)))",
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

func TestParseErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
	}{
		{
			input: "",
		},
		{
			input: "1 +",
		},
		{
			input: "1 + 2 *",
		},
		{
			input: "/ 1",
		},
	}

	for _, test := range tests {
		_, err := NewParser(test.input).Parse()

		if err == nil {
			t.Errorf("expected error for %s, got none", test.input)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	for b.Loop() {
		p := NewParser("1 + 2 * 3 / 4")
		_, _ = p.Parse()
	}
}
