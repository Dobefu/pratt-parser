package errorutil

import (
	"testing"
)

func TestErrorMsg(t *testing.T) {
	t.Parallel()

	err := NewError(ErrorMsgParenNotClosedAtEOF)

	if err.Error() != ErrorMsgParenNotClosedAtEOF {
		t.Errorf(
			"expected error to be '%s', got '%s'",
			ErrorMsgParenNotClosedAtEOF,
			err.Error(),
		)
	}
}
