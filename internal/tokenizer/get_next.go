package tokenizer

import (
	"errors"
)

// GetNext gets the next byte in the expression.
func (t *Tokenizer) GetNext() (byte, error) {
	if t.isEOF {
		return 0, errors.New("cannot get next byte after EOF")
	}

	next := t.exp[t.expIdx]
	t.expIdx++

	if t.expIdx >= t.expLen {
		t.isEOF = true
	}

	return next, nil
}
