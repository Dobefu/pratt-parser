package evaluator

import (
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestEvaluateBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.BinaryExpr
		expected float64
	}{
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5"},
				Right: &ast.NumberLiteral{Value: "5"},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
			},
			expected: 10,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5"},
				Right: &ast.NumberLiteral{Value: "5"},
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
				},
			},
			expected: 0,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5"},
				Right: &ast.NumberLiteral{Value: "5"},
				Operator: token.Token{
					Atom:      "*",
					TokenType: token.TokenTypeOperationMul,
				},
			},
			expected: 25,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5"},
				Right: &ast.NumberLiteral{Value: "5"},
				Operator: token.Token{
					Atom:      "/",
					TokenType: token.TokenTypeOperationDiv,
				},
			},
			expected: 1,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5"},
				Right: &ast.NumberLiteral{Value: "5"},
				Operator: token.Token{
					Atom:      "%",
					TokenType: token.TokenTypeOperationMod,
				},
			},
			expected: 0,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5"},
				Right: &ast.NumberLiteral{Value: "5"},
				Operator: token.Token{
					Atom:      "^",
					TokenType: token.TokenTypeOperationPow,
				},
			},
			expected: 3125,
		},
	}

	for _, test := range tests {
		result, err := NewEvaluator().evaluateBinaryExpr(test.input)

		if err != nil {
			t.Errorf("error evaluating %s: %v", test.input.Expr(), err)
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}

func TestEvaluateBinaryExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input *ast.BinaryExpr
	}{
		{
			input: &ast.BinaryExpr{
				Left:  nil,
				Right: &ast.NumberLiteral{Value: "5"},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
			},
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5"},
				Right: nil,
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
			},
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0"},
				Right: &ast.NumberLiteral{Value: "0"},
				Operator: token.Token{
					Atom:      "/",
					TokenType: token.TokenTypeOperationDiv,
				},
			},
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0"},
				Right: &ast.NumberLiteral{Value: "0"},
				Operator: token.Token{
					Atom:      "%",
					TokenType: token.TokenTypeOperationMod,
				},
			},
		},
	}

	for _, test := range tests {
		result, err := NewEvaluator().evaluateBinaryExpr(test.input)

		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if result != 0 {
			t.Errorf("expected 0, got %f", result)
		}
	}
}
