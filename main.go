// The main entrypoint of the application.
package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/parser"
)

// Main is the main entrypoint of the application.
type Main struct {
	args    []string
	onError func(error)

	ast ast.ExprNode
}

// Run actually runs the application.
func (m *Main) Run() {
	if len(m.args) <= 1 {
		m.onError(errors.New("usage: go run main.go <expression>"))

		return
	}

	p := parser.NewParser(m.args[1])
	ast, err := p.Parse()

	if err != nil {
		m.onError(err)
	}

	m.ast = ast

	if m.ast != nil {
		slog.Info(m.ast.Expr())
	}
}

func main() {
	(&Main{
		args: os.Args,
		onError: func(err error) {
			slog.Error(err.Error())
			os.Exit(1)
		},

		ast: nil,
	}).Run()
}
