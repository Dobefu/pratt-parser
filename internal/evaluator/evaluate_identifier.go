package evaluator

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
)

type identifierInfo struct {
	handler func() (float64, error)
}

func (e *Evaluator) evaluateIdentifier(
	i *ast.Identifier,
) (float64, error) {
	identifier, ok := identifierRegistry[i.Value]

	if !ok {
		return 0, fmt.Errorf("undefined identifier: %s", i.Value)
	}

	return identifier.handler()
}
