package ast

import (
	"github.com/Dobefu/pratt-parser/internal/token"
)

// PrefixExpr defines a struct for a prefix expression.
type PrefixExpr struct {
	Operator token.Token
	Operand  ExprNode
}

// Expr returns the expression of the prefix expression.
func (e *PrefixExpr) Expr() {
	// TODO
}
