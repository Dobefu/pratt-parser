package evaluator

import (
	"math"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
)

func TestEvaluateIdentifier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.Identifier
		expected float64
	}{
		{
			input:    &ast.Identifier{Value: "PI"},
			expected: math.Pi,
		},
	}

	for _, test := range tests {
		result, err := NewEvaluator().evaluateIdentifier(test.input)

		if err != nil {
			t.Errorf("error evaluating %s: %v", test.input.Expr(), err)
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}

func TestEvaluateIdentifierErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.Identifier
		expected string
	}{
		{
			input:    &ast.Identifier{Value: "bogus"},
			expected: "undefined identifier: bogus",
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator().evaluateIdentifier(test.input)

		if err == nil {
			t.Errorf("expected error, got nil")
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
