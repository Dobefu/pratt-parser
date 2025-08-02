package tokenizer

import "testing"

func BenchmarkParseNumber(b *testing.B) {
	t := NewTokenizer("1 + -2 * 3 / 4")

	for b.Loop() {
		t.parseNumber('1')
	}
}
