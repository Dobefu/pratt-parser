package tokenizer

// IsEOF checks if the expression has ended.
func (t *Tokenizer) IsEOF() bool {
	return t.isEOF
}
