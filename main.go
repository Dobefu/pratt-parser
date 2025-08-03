// The main entrypoint of the application.
package main

import (
	"errors"
	"log/slog"
	"os"
	"strconv"

	"github.com/Dobefu/pratt-parser/internal/parser"
)

// Main is the main entrypoint of the application.
type Main struct {
	args    []string
	onError func(error)

	result float64
}

// Run actually runs the application.
func (m *Main) Run() {
	if len(m.args) <= 1 {
		m.onError(errors.New("usage: go run main.go <expression>"))

		return
	}

	p := parser.NewParser(m.args[1])
	result, err := p.Parse()

	if err != nil {
		m.onError(err)
	}

	m.result = result

	slog.Info(strconv.FormatFloat(m.result, 'g', -1, 64))
}

func main() {
	(&Main{
		args: os.Args,
		onError: func(err error) {
			slog.Error(err.Error())
			os.Exit(1)
		},

		result: 0,
	}).Run()
}
