// Package tokenizer provides a struct that handles tokenisation of a string.
package tokenizer

// Tokenizer defines the tokenizer itself.
type Tokenizer struct {
	exp    string
	expLen int
	expIdx int
	isEOF  bool
}

// NewTokenizer creates a new instance of the Tokenizer struct.
func NewTokenizer(exp string) *Tokenizer {
	return &Tokenizer{
		exp:    exp,
		expLen: len(exp),
		expIdx: 0,
		isEOF:  len(exp) <= 0,
	}
}
