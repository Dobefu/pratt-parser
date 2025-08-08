package errorutil

import (
	"errors"
	"testing"
)

func TestErrorMsg(t *testing.T) {
	t.Parallel()

	err := NewError(ErrorMsgParenNotClosedAtEOF)

	if errors.Unwrap(err).Error() != ErrorMsgParenNotClosedAtEOF {
		t.Errorf(
			"expected error to be '%s', got '%s'",
			ErrorMsgParenNotClosedAtEOF,
			errors.Unwrap(err).Error(),
		)
	}
}
