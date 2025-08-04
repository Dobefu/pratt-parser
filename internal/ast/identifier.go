package ast

// Identifier defines a struct for an identifier.
type Identifier struct {
	Value string
}

// Expr returns the expression of the identifier.
func (e *Identifier) Expr() string {
	return e.Value
}
