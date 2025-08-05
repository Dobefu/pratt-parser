package main

import (
	"os"
	"testing"

	"github.com/Dobefu/pratt-parser/internal/errorutil"
)

func TestMain(t *testing.T) {
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
			input:    "1 + 2 * 3",
			expected: 7,
		},
		{
			input:    "(1 + 2) * 3",
			expected: 9,
		},
		{
			input:    "1 + 2 * 3 / 4",
			expected: 2.5,
		},
	}

	for _, test := range tests {
		main := &Main{
			args: []string{os.Args[0], test.input},
			onError: func(err error) {
				t.Errorf("expected no error, got %v", err)
			},

			result: 0,
		}

		main.Run()

		if main.result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, main.result)
		}
	}
}

func TestMainErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "usage: go run main.go <expression>",
		},
		{
			input:    "1 +",
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
	}

	for _, test := range tests {
		var mainErr error
		args := []string{os.Args[0]}

		if test.input != "" {
			args = append(args, test.input)
		}

		main := &Main{
			args: args,
			onError: func(err error) {
				mainErr = err
			},

			result: 0,
		}

		main.Run()

		if mainErr == nil {
			t.Fatalf("expected error, got none for input \"%s\"", test.input)
		}

		if mainErr.Error() != test.expected {
			t.Errorf(
				"expected error \"%s\", got \"%s\"",
				test.expected,
				mainErr.Error(),
			)
		}
	}
}
