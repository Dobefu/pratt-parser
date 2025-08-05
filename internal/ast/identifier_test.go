package ast

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ExprNode
		expected string
	}{
		{
			input:    &Identifier{Value: "PI"},
			expected: "PI",
		},
	}

	for _, test := range tests {
		if test.input.Expr() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, test.input.Expr())
		}
	}
}
