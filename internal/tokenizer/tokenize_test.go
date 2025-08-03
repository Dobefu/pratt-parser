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
			input: "1 + 1",
			expected: []token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
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
			}},
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

func BenchmarkTokenize(b *testing.B) {
	t := NewTokenizer("1 + -2 * 3")

	for b.Loop() {
		_, _ = t.Tokenize()
	}
}
