package parser

import (
	"reflect"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestParseFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []token.Token
		expected *ast.FunctionCall
	}{
		{
			input: []token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.FunctionCall{
				FunctionName: "abs",
				Arguments:    []ast.ExprNode{&ast.NumberLiteral{Value: "1"}},
			},
		},
		{
			input: []token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ",", TokenType: token.TokenTypeComma},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
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

		expr, err := parser.parseFunctionCall("abs", 0)

		if err != nil {
			t.Errorf("expected no error, got \"%v\"", err)
		}

		if !reflect.DeepEqual(expr, test.expected) {
			t.Errorf("expected \"%v\", got \"%v\"", test.expected, expr)
		}
	}
}

func TestParseFunctionCallErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []token.Token
		expected string
	}{
		{
			input:    []token.Token{},
			expected: "cannot get next token after EOF",
		},
		{
			input: []token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
			},
			expected: "cannot peek next token after EOF",
		},
		{
			input: []token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
			},
			expected: "expected '(', got: abs",
		},
		{
			input: []token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "cannot peek next token after EOF",
		},
		{
			input: []token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ",", TokenType: token.TokenTypeComma},
			},
			expected: "cannot get next token after EOF",
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		_, err := parser.parseFunctionCall("abs", 0)

		if err == nil {
			t.Fatalf("expected error, got none for input \"%v\"", test.input)
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
