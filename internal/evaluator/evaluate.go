package evaluator

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/ast"
)

// Evaluate runs the evaluation logic.
func (e *Evaluator) Evaluate(currentAst ast.ExprNode) (float64, error) {
	switch node := currentAst.(type) {
	case *ast.BinaryExpr:
		return e.evaluateBinaryExpr(node)

	case *ast.NumberLiteral:
		return e.evaluateNumberLiteral(node)

	case *ast.PrefixExpr:
		return e.evaluatePrefixExpr(node)

	default:
		return 0, fmt.Errorf("unknown node type: %T", node)
	}
}
