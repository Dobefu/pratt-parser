package parser

import (
	"testing"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []token.Token
		expected string
	}{
		{
			input: []token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "1",
		},
		{
			input: []token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "10", TokenType: token.TokenTypeNumber},
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "5", TokenType: token.TokenTypeNumber},
				{Atom: "**", TokenType: token.TokenTypeOperationPow},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "*", TokenType: token.TokenTypeOperationMul},
				{Atom: "5", TokenType: token.TokenTypeNumber},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "123", TokenType: token.TokenTypeNumber},
				{Atom: "%", TokenType: token.TokenTypeOperationMod},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "5", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: "((1 + (10 / ((5 ** 2) * 5))) - (- (123 % (- 5))))",
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got %v", err)

			continue
		}

		if result.Expr() != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestParseErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []token.Token
		expected string
	}{
		{
			input:    []token.Token{},
			expected: "no tokens to parse",
		},
		{
			input: []token.Token{
				{Atom: "_", TokenType: token.TokenTypeNumber},
			},
			expected: "cannot get next token after EOF",
		},
		{
			input: []token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: "cannot get next token after EOF",
		},
		{
			input: []token.Token{
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "unexpected token: '/'",
		},
	}

	for _, test := range tests {
		p := NewParser(test.input)

		// Test the EOF error.
		if len(test.input) == 1 && test.input[0].Atom == "_" {
			p.isEOF = true
		}

		_, err := p.Parse()

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

func BenchmarkParse(b *testing.B) {
	for b.Loop() {
		p := NewParser(
			[]token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "*", TokenType: token.TokenTypeOperationMul},
				{Atom: "3", TokenType: token.TokenTypeNumber},
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "4", TokenType: token.TokenTypeNumber},
				{Atom: "**", TokenType: token.TokenTypeOperationPow},
				{Atom: "4", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
		)
		_, _ = p.Parse()
	}
}
