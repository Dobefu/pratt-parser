package tokenizer

import (
	"testing"
)

func BenchmarkTokenize(b *testing.B) {
	t := NewTokenizer("1 + -2 * 3")

	for b.Loop() {
		_, _ = t.Tokenize()
	}
}
