package charutil

import (
	"testing"
)

func TestIsLetter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    rune
		expected bool
	}{
		{
			input:    'a',
			expected: true,
		},
		{
			input:    'Z',
			expected: true,
		},
		{
			input:    '0',
			expected: false,
		},
		{
			input:    '9',
			expected: false,
		},
		{
			input:    '🙂',
			expected: false,
		},
	}

	for _, test := range tests {
		if IsLetter(test.input) != test.expected {
			t.Errorf("expected %t, got %t", test.expected, IsLetter(test.input))
		}
	}
}
