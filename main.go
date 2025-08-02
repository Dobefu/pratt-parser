// The main entrypoint of the application.
package main

import (
	"log/slog"
	"os"

	"github.com/Dobefu/pratt-parser/internal/parser"
)

func main() {
	p := parser.NewParser("1 + -2 * 3")
	ast, err := p.Parse()

	slog.Info(ast.Expr())

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
