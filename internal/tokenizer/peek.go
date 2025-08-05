package tokenizer

import (
	"unicode/utf8"

	"github.com/Dobefu/pratt-parser/internal/errorutil"
)

// Peek gets the char of the expression at the current index without advancing it.
func (t *Tokenizer) Peek() (rune, error) {
	if t.isEOF {
		return 0, errorutil.NewError(errorutil.ErrorMsgUnexpectedEOF)
	}

	r, _ := utf8.DecodeRuneInString(t.exp[t.byteIdx:])

	if r == utf8.RuneError {
		return 0, errorutil.NewError(errorutil.ErrorMsgInvalidUTF8Char)
	}

	return r, nil
}
