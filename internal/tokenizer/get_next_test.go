package tokenizer

import "testing"

func TestGetNext(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected rune
	}{
		{
			input:    "1 + 1",
			expected: '1',
		},
		{
			input:    "🙂",
			expected: '🙂',
		},
	}

	for _, test := range tests {
		tokenizer := NewTokenizer(test.input)
		token, err := tokenizer.GetNext()

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if token != test.expected {
			t.Errorf("expected %v, got %v", test.expected, token)
		}
	}
}

func TestGetNextErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
	}{
		{
			input: "\xFF",
		},
	}

	for _, test := range tests {
		tokenizer := NewTokenizer(test.input)
		_, err := tokenizer.GetNext()

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	}
}
