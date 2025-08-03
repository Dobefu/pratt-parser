package tokenizer

import (
	"reflect"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestParseNumber(t *testing.T) {
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
			input: "1.1",
			expected: []token.Token{
				{Atom: "1.1", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1_000_000",
			expected: []token.Token{
				{Atom: "1000000", TokenType: token.TokenTypeNumber},
			},
		},
		{
			input: "1.1_000_000",
			expected: []token.Token{
				{Atom: "1.1000000", TokenType: token.TokenTypeNumber},
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

func BenchmarkParseNumber(b *testing.B) {
	for b.Loop() {
		t := NewTokenizer("1 + -2 * 3 / 4")

		_, _ = t.parseNumber('1')
	}
}
