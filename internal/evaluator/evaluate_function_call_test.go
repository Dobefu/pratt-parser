package evaluator

import (
	"fmt"
	"math"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
)

func evaluateFunctionCallCreateFunctionCall(
	functionName string,
	arguments ...ast.ExprNode,
) ast.ExprNode {
	return &ast.FunctionCall{
		FunctionName: functionName,
		Arguments:    arguments,
	}
}

func TestEvaluateFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected float64
	}{
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"abs",
				&ast.NumberLiteral{Value: "5"},
			),
			expected: 5,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"sin",
				&ast.NumberLiteral{Value: fmt.Sprintf("%f", math.Pi/2)},
			),
			expected: 1,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"cos",
				&ast.NumberLiteral{Value: fmt.Sprintf("%f", math.Pi)},
			),
			expected: -1,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"tan",
				&ast.NumberLiteral{Value: fmt.Sprintf("%f", math.Pi/4)},
			),
			expected: 1,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"sqrt",
				&ast.NumberLiteral{Value: "16"},
			),
			expected: 4,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"round",
				&ast.NumberLiteral{Value: "3.14"},
			),
			expected: 3,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"floor",
				&ast.NumberLiteral{Value: "6.9"},
			),
			expected: 6,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"ceil",
				&ast.NumberLiteral{Value: "3.14"},
			),
			expected: 4,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"min",
				&ast.NumberLiteral{Value: "1"},
				&ast.NumberLiteral{Value: "2"},
			),
			expected: 1,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"max",
				&ast.NumberLiteral{Value: "1"},
				&ast.NumberLiteral{Value: "2"},
			),
			expected: 2,
		},
	}

	for _, test := range tests {
		result, err := NewEvaluator().Evaluate(test.input)
		result = math.Round(result*1000) / 1000

		if err != nil {
			t.Errorf("error evaluating %s: %v", test.input, err)
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}

func TestEvaluateFunctionCallErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected string
	}{
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"bogus",
				&ast.NumberLiteral{Value: "1"},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"abs",
				&ast.NumberLiteral{Value: "1"},
				&ast.NumberLiteral{Value: "1"},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "abs", 1, 2),
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"abs",
				&ast.NumberLiteral{Value: "a"},
			),
			expected: `strconv.ParseFloat: parsing "a": invalid syntax`,
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator().Evaluate(test.input)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != test.expected {
			t.Errorf(
				"expected error \"%v\", got \"%v\"",
				test.expected,
				err.Error(),
			)
		}
	}
}
