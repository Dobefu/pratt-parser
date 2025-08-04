package evaluator

import (
	"math"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestEvaluate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected float64
	}{
		{
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{Value: "5"},
				Right: &ast.BinaryExpr{
					Left: &ast.NumberLiteral{Value: "5"},
					Right: &ast.PrefixExpr{
						Operator: token.Token{
							Atom:      "-",
							TokenType: token.TokenTypeOperationSub,
						},
						Operand: &ast.FunctionCall{
							FunctionName: "abs",
							Arguments: []ast.ExprNode{
								&ast.Identifier{Value: "PI"},
							},
						},
					},
					Operator: token.Token{
						Atom:      "+",
						TokenType: token.TokenTypeOperationAdd,
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
			},
			expected: 5 + math.Abs(-5+math.Pi),
		},
	}

	for _, test := range tests {
		result, err := NewEvaluator().Evaluate(test.input)

		if err != nil {
			t.Errorf("error evaluating %s: %v", test.input, err)
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}
