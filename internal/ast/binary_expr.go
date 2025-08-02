package ast

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/token"
)

// BinaryExpr defines a struct for a binary expression.
type BinaryExpr struct {
	Left     ExprNode
	Right    ExprNode
	Operator token.Token
}

// Expr returns the expression of the binary expression.
func (e *BinaryExpr) Expr() string {
	return fmt.Sprintf("(%s %s %s)", e.Left.Expr(), e.Operator.Atom, e.Right.Expr())
}
