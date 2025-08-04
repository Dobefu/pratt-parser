package tokenizer

import (
	"errors"
	"unicode/utf8"
)

// Peek gets the char of the expression at the current index without advancing it.
func (t *Tokenizer) Peek() (rune, error) {
	if t.isEOF {
		return 0, errors.New("cannot peek next character after EOF")
	}

	r, _ := utf8.DecodeRuneInString(t.exp[t.byteIdx:])

	if r == utf8.RuneError {
		return 0, errors.New("invalid UTF-8 sequence")
	}

	return r, nil
}
