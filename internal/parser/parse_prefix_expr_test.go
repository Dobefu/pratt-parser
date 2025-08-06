package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestParsePrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []*token.Token
		expected *ast.PrefixExpr
	}{
		{
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				Operand:  &ast.NumberLiteral{Value: "1"},
			},
		},
		{
			input: []*token.Token{
				{Atom: "-", TokenType: token.TokenTypeOperationAdd},
				{Atom: "PI", TokenType: token.TokenTypeIdentifier},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{Atom: "-", TokenType: token.TokenTypeOperationSub},
				Operand:  &ast.NumberLiteral{Value: "PI"},
			},
		},
		{
			input: []*token.Token{
				{Atom: "-", TokenType: token.TokenTypeOperationAdd},
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{Atom: "-", TokenType: token.TokenTypeOperationSub},
				Operand:  &ast.NumberLiteral{Value: "abs(1)"},
			},
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got \"%s\"", err.Error())

			continue
		}

		if result.Expr() != test.expected.Expr() {
			t.Errorf("expected \"%s\", got \"%s\"", test.expected.Expr(), result.Expr())
		}
	}
}

func TestParsePrefixExprErr(t *testing.T) {
	t.Parallel()

	errNextTokenAfterEOF := errorutil.ErrorMsgUnexpectedEOF

	tests := []struct {
		input    []*token.Token
		expected string
	}{
		{
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: errNextTokenAfterEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
			},
			expected: errNextTokenAfterEOF,
		},
		{
			input: []*token.Token{
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, ")"),
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: errNextTokenAfterEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "(", TokenType: token.TokenTypeLParen},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgExpectedCloseParen, "("),
		},
	}

	for _, test := range tests {
		_, err := NewParser(test.input).Parse()

		if err == nil {
			t.Fatalf("expected error for \"%v\", got none", test.input)
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
