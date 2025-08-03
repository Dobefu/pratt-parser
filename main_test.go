package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1",
			expected: "1",
		},
		{
			input:    "1 + 1",
			expected: "(1 + 1)",
		},
		{
			input:    "(1 + 1)",
			expected: "(1 + 1)",
		},
		{
			input:    "1 + 2 * 3",
			expected: "(1 + (2 * 3))",
		},
		{
			input:    "(1 + 2) * 3",
			expected: "((1 + 2) * 3)",
		},
		{
			input:    "1 + 2 * 3 / 4",
			expected: "(1 + ((2 * 3) / 4))",
		},
	}

	for _, test := range tests {
		main := &Main{
			args: []string{os.Args[0], test.input},
			onError: func(err error) {
				t.Errorf("expected no error, got %v", err)
			},

			ast: nil,
		}

		main.Run()

		if main.ast.Expr() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, main.ast.Expr())
		}
	}
}

func TestMainErr(t *testing.T) {
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
	}

	for _, test := range tests {
		hasError := false
		args := []string{os.Args[0]}

		if test.input != "" {
			args = append(args, test.input)
		}

		main := &Main{
			args: args,
			onError: func(err error) {
				hasError = true
			},

			ast: nil,
		}

		main.Run()

		if !hasError {
			t.Errorf("expected error, got none for input %s", test.input)
		}
	}
}
