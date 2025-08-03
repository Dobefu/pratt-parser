package tokenizer

import (
	"reflect"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestTokenize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected []token.Token
	}{
		{
			input: "1",
			expected: []token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1e0",
			expected: []token.Token{
				{Atom: "1e0", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1e6",
			expected: []token.Token{
				{Atom: "1e6", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1e+6",
			expected: []token.Token{
				{Atom: "1e6", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1E6",
			expected: []token.Token{
				{Atom: "1E6", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1e-6",
			expected: []token.Token{
				{Atom: "1e-6", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1.2e6",
			expected: []token.Token{
				{Atom: "1.2e6", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1 + 1",
			expected: []token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1 ** 1",
			expected: []token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "**", TokenType: token.TokenTypeOperationPow},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "10 % 3",
			expected: []token.Token{
				{Atom: "10", TokenType: token.TokenTypeNumber},
				{Atom: "%", TokenType: token.TokenTypeOperationMod},
				{Atom: "3", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1 + 2 * 3 / 4",
			expected: []token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "*", TokenType: token.TokenTypeOperationMul},
				{Atom: "3", TokenType: token.TokenTypeNumber},
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "4", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1 + 2 * 3 / 4 - 5",
			expected: []token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "*", TokenType: token.TokenTypeOperationMul},
				{Atom: "3", TokenType: token.TokenTypeNumber},
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "4", TokenType: token.TokenTypeNumber},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "5", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "()",
			expected: []token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
		},
	}

	for _, test := range tests {
		tokens, err := NewTokenizer(test.input).Tokenize()

		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tokens, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, tokens)
		}
	}
}

func TestTokenizeErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
	}{
		{
			input: "1e",
		},
		{
			input: "1e-",
		},
		{
			input: "1e-r",
		},
		{
			input: "1e6e6",
		},
		{
			input: "1e6er",
		},
		{
			input: "💔",
		},
		{
			input: "*",
		},
	}

	for _, test := range tests {
		_, err := NewTokenizer(test.input).Tokenize()

		if err == nil {
			t.Fatalf("expected error, got none for input %s", test.input)
		}
	}
}

func BenchmarkTokenize(b *testing.B) {
	t := NewTokenizer("1 + -2 * 3")

	for b.Loop() {
		_, _ = t.Tokenize()
	}
}
