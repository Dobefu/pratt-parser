package ast

import (
	"fmt"

	"github.com/Dobefu/pratt-parser/internal/token"
)

// PrefixExpr defines a struct for a prefix expression.
type PrefixExpr struct {
	Operator token.Token
	Operand  ExprNode
}

// Expr returns the expression of the prefix expression.
func (e *PrefixExpr) Expr() string {
	return fmt.Sprintf("(%s %s)", e.Operator.Atom, e.Operand.Expr())
}
