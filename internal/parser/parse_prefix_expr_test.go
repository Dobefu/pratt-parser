package parser

import (
	"math"
	"testing"
)

func TestParsePrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected float64
	}{
		{
			input:    "1",
			expected: 1,
		},
		{
			input:    "PI",
			expected: math.Pi,
		},
		{
			input:    "PI + 1",
			expected: math.Pi + 1,
		},
		{
			input:    "abs(-1)",
			expected: 1,
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got %v", err)

			continue
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
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
