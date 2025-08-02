package parser

import "testing"

func BenchmarkParse(b *testing.B) {
	for b.Loop() {
		p := NewParser("1 + 2 * 3 / 4")
		_, _ = p.Parse()
	}
}
