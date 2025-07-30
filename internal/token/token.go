// Package token defines a Token struct.
package token

// Token defines a single expression token.
type Token struct {
	Atom      string
	TokenType Type
}
