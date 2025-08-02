package token

// Type defines an enum of possible token types.
type Type int

const (
	// TokenTypeOperationAdd represents the addition operation.
	TokenTypeOperationAdd = iota
	// TokenTypeOperationSub represents the subtraction operation.
	TokenTypeOperationSub
	// TokenTypeOperationMul represents the multiplication operation.
	TokenTypeOperationMul
	// TokenTypeOperationDiv represents the division operation.
	TokenTypeOperationDiv
	// TokenTypeNumber represents a number literal.
	TokenTypeNumber
	// TokenTypeIdentifier represents an identifier.
	TokenTypeIdentifier
	// TokenTypeLParen represents a left parenthesis.
	TokenTypeLParen
	// TokenTypeRParen represents a right parenthesis.
	TokenTypeRParen
)
