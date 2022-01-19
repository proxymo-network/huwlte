package huwlte

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorCode_String(t *testing.T) {
	for _, test := range []struct {
		Code ErrorCode
		Str  string
	}{
		{ErrorCodeInternal, "internal"},
		{ErrorCodeNotSupported, "not supported"},
		{ErrorCodeNoRights, "no rights"},
		{ErrorCodeBusy, "busy"},
		{ErrorCodeCSRF, "csrf"},
		{ErrorCode(123), "unknown(123)"},
	} {
		assert.Equal(t, test.Str, test.Code.String())
	}
}

func TestError_Error(t *testing.T) {
	for _, test := range []struct {
		Err *Error
		Str string
	}{
		{
			&Error{Code: ErrorCodeInternal, Message: "shit happens"},
			"100001: internal (shit happens)",
		},
		{
			&Error{Code: ErrorCodeNotSupported},
			"100002: not supported",
		},
	} {
		assert.Equal(t, test.Str, test.Err.Error())
	}
}

func TestIsError(t *testing.T) {
	err := newError(int(ErrorCodeInternal), "shit happens")

	assert.True(t, IsError(err, ErrorCodeInternal))
	assert.False(t, IsError(err, ErrorCodeNotSupported))
	assert.False(t, IsError(context.Canceled, ErrorCodeInternal))
}
