package evaluator

import (
	"fmt"
	"math"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
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
		{
			input:    &ast.Identifier{Value: "TAU"},
			expected: math.Pi * 2,
		},
		{
			input:    &ast.Identifier{Value: "E"},
			expected: math.E,
		},
		{
			input:    &ast.Identifier{Value: "PHI"},
			expected: math.Phi,
		},
		{
			input:    &ast.Identifier{Value: "LN2"},
			expected: math.Ln2,
		},
		{
			input:    &ast.Identifier{Value: "LN10"},
			expected: math.Ln10,
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
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"),
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator().evaluateIdentifier(test.input)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != test.expected {
			t.Errorf(
				"expected error \"%s\", got \"%s\"",
				test.expected,
				err.Error(),
			)
		}
	}
}
