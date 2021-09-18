package huwlte

import "fmt"

type ErrorCode int

// String returns the string representation of the error code.
func (ec ErrorCode) String() string {
	switch ec {
	case ErrorCodeInternal:
		return "internal"
	case ErrorCodeNotSupported:
		return "not supported"
	case ErrorCodeNoRights:
		return "no rights"
	case ErrorCodeBusy:
		return "busy"
	case ErrorCodeCSRF:
		return "csrf"
	default:
		return fmt.Sprintf("unknown(%d)", ec)
	}
}

const (
	ErrorCodeInternal     ErrorCode = 100001
	ErrorCodeNotSupported ErrorCode = 100002
	ErrorCodeNoRights     ErrorCode = 100003
	ErrorCodeBusy         ErrorCode = 100004
	ErrorCodeCSRF         ErrorCode = 125002
)

type Error struct {
	Code    ErrorCode
	Message string
}

func newError(ec int, message string) *Error {
	return &Error{
		Code:    ErrorCode(ec),
		Message: message,
	}
}

func (err *Error) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("%d: %s", err.Code, err.Code.String())
	}
	return fmt.Sprintf("%d: %s (%s)", err.Code, err.Code.String(), err.Message)
}
