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
				Left:  &NumberLiteral{Value: "1"},
				Right: &NumberLiteral{Value: "1"},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
			},
			expected: "",
		},
	}

	for _, test := range tests {
		if test.input.Expr() != "(1 + 1)" {
			t.Errorf("expected %s, got %s", test.expected, test.input.Expr())
		}
	}
}
