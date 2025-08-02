package parser

import (
	"testing"

	"github.com/Dobefu/pratt-parser/internal/token"
)

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
