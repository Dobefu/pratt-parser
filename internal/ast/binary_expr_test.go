package ast

import (
	"testing"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ExprNode
		expected string
	}{
		{
			input: &BinaryExpr{
				Left:  &NumberLiteral{Value: "1", Pos: 0},
				Right: &NumberLiteral{Value: "1", Pos: 2},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				Pos: 0,
			},
			expected: "(1 + 1)",
		},
	}

	for _, test := range tests {
		if test.input.Expr() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, test.input.Expr())
		}
	}
}
