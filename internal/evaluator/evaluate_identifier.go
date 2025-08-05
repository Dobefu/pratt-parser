package evaluator

import (
	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
)

type identifierInfo struct {
	handler func() (float64, error)
}

func (e *Evaluator) evaluateIdentifier(
	i *ast.Identifier,
) (float64, error) {
	identifier, ok := identifierRegistry[i.Value]

	if !ok {
		return 0, errorutil.NewError(errorutil.ErrorMsgUndefinedIdentifier, i.Value)
	}

	return identifier.handler()
}
