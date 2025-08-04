package tokenizer

import (
	"errors"
	"unicode/utf8"
)

// GetNext gets the next character in the expression.
func (t *Tokenizer) GetNext() (rune, error) {
	if t.isEOF {
		return 0, errors.New("cannot get next character after EOF")
	}

	r, size := utf8.DecodeRuneInString(t.exp[t.byteIdx:])

	if r == utf8.RuneError {
		return 0, errors.New("invalid UTF-8 sequence")
	}

	t.byteIdx += size
	t.expIdx++

	if t.expIdx >= t.expLen {
		t.isEOF = true
	}

	return r, nil
}
