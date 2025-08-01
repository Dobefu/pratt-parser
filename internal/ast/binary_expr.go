package ast

import (
	"github.com/Dobefu/pratt-parser/internal/token"
)

// BinaryExpr defines a struct for a binary expression.
type BinaryExpr struct {
	Left     ExprNode
	Right    ExprNode
	Operator token.Token
}

// Expr returns the expression of the binary expression.
func (e *BinaryExpr) Expr() {
	// TODO
}
