package ast

// FunctionCall defines a struct for a function call.
type FunctionCall struct {
	FunctionName string
	Arguments    []ExprNode
}

// Expr returns the expression of the function call.
func (fc *FunctionCall) Expr() string {
	args := make([]string, len(fc.Arguments))

	for i, arg := range fc.Arguments {
		args[i] = arg.Expr()
	}

	return fc.FunctionName + "(" + join(args, ", ") + ")"
}

func join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	result := strs[0]

	for _, s := range strs[1:] {
		result += sep + s
	}

	return result
}
