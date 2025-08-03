package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected float64
	}{
		{
			input:    "1",
			expected: 1,
		},
		{
			input:    "1 + 1",
			expected: 2,
		},
		{
			input:    "(1 + 1)",
			expected: 2,
		},
		{
			input:    "1 ** 1",
			expected: 1,
		},
		{
			input:    "1 % 1",
			expected: 0,
		},
		{
			input:    "1 + 2 * 3",
			expected: 7,
		},
		{
			input:    "1 + 2 * 3 / 4",
			expected: 2.5,
		},
		{
			input:    "1 + 2 * 3 / 4 - 5",
			expected: -2.5,
		},
		{
			input:    "(1 + 10 / 5 + 2 * 5 - ( -123 - (-(5))))",
			expected: 131,
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

func TestParseErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
	}{
		{
			input: "",
		},
		{
			input: "1 +",
		},
		{
			input: "1 + 2 *",
		},
		{
			input: "/ 1",
		},
	}

	for _, test := range tests {
		_, err := NewParser(test.input).Parse()

		if err == nil {
			t.Errorf("expected error for %s, got none", test.input)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	for b.Loop() {
		p := NewParser("1 + 2 * 3 / 4")
		_, _ = p.Parse()
	}
}
