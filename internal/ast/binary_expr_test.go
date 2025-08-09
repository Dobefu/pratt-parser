package ast

import (
	"testing"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input       ExprNode
		expected    string
		expectedPos int
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
			expected:    "(1 + 1)",
			expectedPos: 0,
		},
		{
			input: &BinaryExpr{
				Left:  &NumberLiteral{Value: "1", Pos: 0},
				Right: &NumberLiteral{Value: "2", Pos: 2},
				Operator: token.Token{
					Atom:      "*",
					TokenType: token.TokenTypeOperationMul,
				},
				Pos: 0,
			},
			expected:    "(1 * 2)",
			expectedPos: 0,
		},
	}

	for _, test := range tests {
		if test.input.Expr() != test.expected {
			t.Errorf("expected '%s', got '%s'", test.expected, test.input.Expr())
		}

		if test.input.Position() != test.expectedPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedPos, test.input.Position())
		}
	}
}
