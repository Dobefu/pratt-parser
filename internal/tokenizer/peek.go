package tokenizer

import "errors"

// Peek gets the byte of the expression at the current index without advancing it.
func (t *Tokenizer) Peek() (byte, error) {
	if t.isEOF {
		return 0, errors.New("cannot get next byte after EOF")
	}

	return t.exp[t.expIdx], nil
}
