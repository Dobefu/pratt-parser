package parser

import (
	"reflect"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
)

func TestParseFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected ast.ExprNode
	}{
		{
			input: "(1)",
			expected: &ast.FunctionCall{
				FunctionName: "abs",
				Arguments:    []ast.ExprNode{&ast.NumberLiteral{Value: "1"}},
			},
		},
		{
			input: "(1, 1)",
			expected: &ast.FunctionCall{
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1"},
					&ast.NumberLiteral{Value: "1"},
				},
			},
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		tokens, err := parser.tokenizer.Tokenize()

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		parser.tokens = tokens
		parser.tokenLen = len(tokens)

		expr, err := parser.parseFunctionCall("abs", 0)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if !reflect.DeepEqual(expr, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, expr)
		}
	}
}

func TestParseFunctionCallErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
	}{
		{
			input: "",
		},
		{
			input: "(",
		},
		{
			input: "abs",
		},
		{
			input: "abs",
		},
		{
			input: "(1",
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		tokens, err := parser.tokenizer.Tokenize()

		if err != nil {
			t.Errorf("expected a different error, got %v", err)
		}

		parser.tokens = tokens
		parser.tokenLen = len(tokens)

		_, err = parser.parseFunctionCall("abs", 0)

		if err == nil {
			t.Errorf("expected error, got none for input %s", test.input)
		}
	}
}
