package parser

import (
	"testing"

	"github.com/Dobefu/pratt-parser/internal/token"
)

func TestParseBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected float64
	}{
		{
			input:    "1 + 1",
			expected: 2,
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got %v", err)

			continue
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}

func TestParseBinaryExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "no tokens to parse",
		},
		{
			input:    "1 +",
			expected: "cannot get next token after EOF",
		},
		{
			input:    "1 + 2 *",
			expected: "cannot peek next character after EOF",
		},
		{
			input:    "1 + 2 * 3 /",
			expected: "cannot get next token after EOF",
		},
		{
			input:    "1 💔 1",
			expected: "unexpected character: 💔 at position 3",
		},
	}

	for _, test := range tests {
		_, err := NewParser(test.input).Parse()

		if err == nil {
			t.Errorf("expected error for %s, got none", test.input)
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

func BenchmarkParseBinaryExpr(b *testing.B) {
	for b.Loop() {
		p := NewParser("")

		_, _ = p.parseBinaryExpr(
			&token.Token{
				Atom:      "1",
				TokenType: token.TokenTypeOperationAdd,
			},
			nil,
			nil,
			0,
		)
	}
}
