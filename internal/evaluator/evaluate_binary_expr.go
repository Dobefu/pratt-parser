package evaluator

import (
	"fmt"
	"math"

	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/token"
)

func (e *Evaluator) evaluateBinaryExpr(node *ast.BinaryExpr) (float64, error) {
	leftEvaluated, err := e.Evaluate(node.Left)

	if err != nil {
		return 0, err
	}

	rightEvaluated, err := e.Evaluate(node.Right)

	if err != nil {
		return 0, err
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return leftEvaluated + rightEvaluated, nil

	case token.TokenTypeOperationSub:
		return leftEvaluated - rightEvaluated, nil

	case token.TokenTypeOperationMul:
		return leftEvaluated * rightEvaluated, nil

	case token.TokenTypeOperationDiv:
		return leftEvaluated / rightEvaluated, nil

	case token.TokenTypeOperationMod:
		return math.Mod(leftEvaluated, rightEvaluated), nil

	case token.TokenTypeOperationPow:
		return math.Pow(leftEvaluated, rightEvaluated), nil

	default:
		return 0, fmt.Errorf("unknown operator: %s", node.Operator.Atom)
	}
}
