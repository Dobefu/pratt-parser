package evaluator

import (
	"github.com/Dobefu/pratt-parser/internal/ast"
	"github.com/Dobefu/pratt-parser/internal/errorutil"
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

	case *ast.FunctionCall:
		return e.evaluateFunctionCall(node)

	case *ast.Identifier:
		return e.evaluateIdentifier(node)

	default:
		return 0, errorutil.NewError(errorutil.ErrorMsgUnknownNodeType, node)
	}
}
