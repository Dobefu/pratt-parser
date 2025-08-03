// The main entrypoint of the application.
package main

import (
	"log/slog"
	"os"

	"github.com/Dobefu/pratt-parser/internal/parser"
)

func main() {
	if len(os.Args) <= 1 {
		slog.Error("Usage: go run main.go <expression>")
		os.Exit(1)
	}

	p := parser.NewParser(os.Args[1])
	ast, err := p.Parse()

	if ast != nil {
		slog.Info(ast.Expr())
	}

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
