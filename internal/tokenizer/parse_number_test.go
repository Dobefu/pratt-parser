package tokenizer

import (
	"reflect"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func getNumberToken(atom string) token.Token {
	return token.Token{
		Atom:      atom,
		TokenType: token.TokenTypeNumber,
	}
}

func TestParseNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected []token.Token
	}{
		{
			input:    "1",
			expected: []token.Token{getNumberToken("1")},
		},
		{
			input:    "1.1",
			expected: []token.Token{getNumberToken("1.1")},
		},
		{
			input:    "1e1",
			expected: []token.Token{getNumberToken("1e1")},
		},
		{
			input:    "1e+1",
			expected: []token.Token{getNumberToken("1e1")},
		},
		{
			input:    "1e-1",
			expected: []token.Token{getNumberToken("1e-1")},
		},
		{
			input: "1+1",
			expected: []token.Token{
				getNumberToken("1"),
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				getNumberToken("1"),
			},
		},
		{
			input:    "1_000_000",
			expected: []token.Token{getNumberToken("1000000")},
		},
		{
			input:    "1.1_000_000",
			expected: []token.Token{getNumberToken("1.1000000")},
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

func TestParseNumberErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
	}{
		{
			input: "1__2",
		},
		{
			input: "1..1",
		},
		{
			input: "1.",
		},
	}

	for _, test := range tests {
		_, err := NewTokenizer(test.input).Tokenize()

		if err == nil {
			t.Fatalf("expected error for %s, got none", test.input)
		}
	}
}

func BenchmarkParseNumber(b *testing.B) {
	for b.Loop() {
		t := NewTokenizer("1 + -2 * 3 / 4")

		_, _ = t.parseNumber('1')
	}
}
