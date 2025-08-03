package evaluator

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
)

type functionHandler func([]float64) (float64, error)

type functionInfo struct {
	handler  functionHandler
	argCount int
}

func (e *Evaluator) evaluateFunctionCall(
	fc *ast.FunctionCall,
) (float64, error) {
	function, ok := functionRegistry[fc.FunctionName]

	if !ok {
		return 0, fmt.Errorf("undefined function: %s", fc.FunctionName)
	}

	argValues, err := e.evaluateArguments(
		fc.Arguments,
		function.argCount,
		fc.FunctionName,
	)

	if err != nil {
		return 0, err
	}

	return function.handler(argValues)
}

func (e *Evaluator) evaluateArguments(
	args []ast.ExprNode,
	expectedCount int,
	functionName string,
) ([]float64, error) {
	argValues := make([]float64, len(args))

	for i, arg := range args {
		val, err := e.Evaluate(arg)

		if err != nil {
			return nil, err
		}

		argValues[i] = val
	}

	if expectedCount > 0 && len(argValues) != expectedCount {
		return nil, fmt.Errorf(
			"%s() expects exactly %d argument(s), but got %d",
			functionName,
			expectedCount,
			len(argValues),
		)
	}

	return argValues, nil
}
