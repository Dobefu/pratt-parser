package evaluator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
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
		input    *ast.PrefixExpr
		expected string
	}{
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
				},
				Operand: nil,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnknownNodeType, nil),
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator().evaluatePrefixExpr(test.input)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if errors.Unwrap(err).Error() != test.expected {
			t.Errorf(
				"expected error \"%s\", got \"%s\"",
				test.expected,
				errors.Unwrap(err).Error(),
			)
		}
	}
}
