package ast

import (
	"testing"
)

func TestFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ExprNode
		expected string
	}{
		{
			input: &FunctionCall{
				FunctionName: "abs",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "1", Pos: 0},
				},
				Pos: 0,
			},
			expected: "abs(1)",
		},
		{
			input: &FunctionCall{
				FunctionName: "max",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "1", Pos: 0},
					&NumberLiteral{Value: "2", Pos: 2},
				},
				Pos: 0,
			},
			expected: "max(1, 2)",
		},
	}

	for _, test := range tests {
		if test.input.Expr() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, test.input.Expr())
		}
	}
}
