package evaluator

import (
	"fmt"
	"math"

	"github.com/Dobefu/pratt-parser/internal/ast"
)

func (e *Evaluator) evaluateFunctionCall(fc *ast.FunctionCall) (float64, error) {
	switch fc.FunctionName {
	case "abs":
		argValues, err := e.evaluateArguments(fc.Arguments, 1, fc.FunctionName)

		if err != nil {
			return 0, err
		}

		return math.Abs(argValues[0]), nil

	case "sin":
		argValues, err := e.evaluateArguments(fc.Arguments, 1, fc.FunctionName)

		if err != nil {
			return 0, err
		}

		return math.Sin(argValues[0]), nil

	case "cos":
		argValues, err := e.evaluateArguments(fc.Arguments, 1, fc.FunctionName)

		if err != nil {
			return 0, err
		}

		return math.Cos(argValues[0]), nil

	case "tan":
		argValues, err := e.evaluateArguments(fc.Arguments, 1, fc.FunctionName)

		if err != nil {
			return 0, err
		}

		return math.Tan(argValues[0]), nil

	case "sqrt":
		argValues, err := e.evaluateArguments(fc.Arguments, 1, fc.FunctionName)

		if err != nil {
			return 0, err
		}

		return math.Sqrt(argValues[0]), nil

	case "round":
		argValues, err := e.evaluateArguments(fc.Arguments, 1, fc.FunctionName)

		if err != nil {
			return 0, err
		}

		return math.Round(argValues[0]), nil

	case "floor":
		argValues, err := e.evaluateArguments(fc.Arguments, 1, fc.FunctionName)

		if err != nil {
			return 0, err
		}

		return math.Floor(argValues[0]), nil

	case "ceil":
		argValues, err := e.evaluateArguments(fc.Arguments, 1, fc.FunctionName)

		if err != nil {
			return 0, err
		}

		return math.Ceil(argValues[0]), nil

	default:
		return 0, fmt.Errorf("undefined function: %s", fc.FunctionName)
	}
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
