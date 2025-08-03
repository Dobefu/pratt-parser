package evaluator

import (
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestEvaluatePrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.PrefixExpr
		expected float64
	}{
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
				},
				Operand: &ast.NumberLiteral{Value: "5"},
			},
			expected: -5,
		},
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				Operand: &ast.NumberLiteral{Value: "5"},
			},
			expected: 5,
		},
	}

	for _, test := range tests {
		result, err := NewEvaluator().evaluatePrefixExpr(test.input)

		if err != nil {
			t.Errorf("error evaluating %s: %v", test.input.Expr(), err)
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}

func TestEvaluatePrefixExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input *ast.PrefixExpr
	}{
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
				},
				Operand: nil,
			},
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator().evaluatePrefixExpr(test.input)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	}
}
