package evaluator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
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
				Left:  &ast.NumberLiteral{Value: "5", Pos: 0},
				Right: &ast.NumberLiteral{Value: "5", Pos: 2},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				Pos: 1,
			},
			expected: 10,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", Pos: 0},
				Right: &ast.NumberLiteral{Value: "5", Pos: 2},
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
				},
				Pos: 1,
			},
			expected: 0,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", Pos: 0},
				Right: &ast.NumberLiteral{Value: "5", Pos: 2},
				Operator: token.Token{
					Atom:      "*",
					TokenType: token.TokenTypeOperationMul,
				},
				Pos: 1,
			},
			expected: 25,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", Pos: 0},
				Right: &ast.NumberLiteral{Value: "5", Pos: 2},
				Operator: token.Token{
					Atom:      "/",
					TokenType: token.TokenTypeOperationDiv,
				},
				Pos: 1,
			},
			expected: 1,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", Pos: 0},
				Right: &ast.NumberLiteral{Value: "5", Pos: 2},
				Operator: token.Token{
					Atom:      "%",
					TokenType: token.TokenTypeOperationMod,
				},
				Pos: 1,
			},
			expected: 0,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", Pos: 0},
				Right: &ast.NumberLiteral{Value: "5", Pos: 2},
				Operator: token.Token{
					Atom:      "^",
					TokenType: token.TokenTypeOperationPow,
				},
				Pos: 1,
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
		input    *ast.BinaryExpr
		expected string
	}{
		{
			input: &ast.BinaryExpr{
				Left:  nil,
				Right: &ast.NumberLiteral{Value: "5", Pos: 2},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				Pos: 1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnknownNodeType, nil),
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", Pos: 0},
				Right: nil,
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				Pos: 1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnknownNodeType, nil),
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0", Pos: 0},
				Right: &ast.NumberLiteral{Value: "0", Pos: 2},
				Operator: token.Token{
					Atom:      "/",
					TokenType: token.TokenTypeOperationDiv,
				},
				Pos: 1,
			},
			expected: errorutil.ErrorMsgDivByZero,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0", Pos: 0},
				Right: &ast.NumberLiteral{Value: "0", Pos: 2},
				Operator: token.Token{
					Atom:      "%",
					TokenType: token.TokenTypeOperationMod,
				},
				Pos: 1,
			},
			expected: errorutil.ErrorMsgModByZero,
		},
		{
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0", Pos: 0},
				Right: &ast.NumberLiteral{Value: "0", Pos: 2},
				Operator: token.Token{
					Atom:      ",",
					TokenType: token.TokenTypeComma,
				},
				Pos: 1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnknownOperator, ","),
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator().evaluateBinaryExpr(test.input)

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
