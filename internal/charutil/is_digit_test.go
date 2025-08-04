package charutil

import "testing"

func TestIsDigit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    rune
		expected bool
	}{
		{
			input:    '0',
			expected: true,
		},
		{
			input:    '9',
			expected: true,
		},
		{
			input:    'a',
			expected: false,
		},
		{
			input:    'z',
			expected: false,
		},
		{
			input:    '🙂',
			expected: false,
		},
	}

	for _, test := range tests {
		if IsDigit(test.input) != test.expected {
			t.Errorf("expected %t, got %t", test.expected, IsDigit(test.input))
		}
	}
}
