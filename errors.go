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
	case ErrorCodeWrongToken:
		return "wrong token"
	case ErrorCodeWrongSessionToken:
		return "wrong session token"
	case ErrorCodeUsernameWrong:
		return "wrong username"
	case ErrorCodePasswordWrong:
		return "wrong password"
	case ErrorCodeAlreadyLogin:
		return "already login"
	case ErrorCodeUsernamePwdWrong:
		return "username or password wrong"
	case ErrorCodeUsernamePwdOrerrun:
		return "username or password orerrun"
	case ErrorCodeUsernamePwdModify:
		return "username or password modify"
	default:
		return fmt.Sprintf("unknown(%d)", ec)
	}
}

const (
	ErrorCodeInternal     ErrorCode = 100001
	ErrorCodeNotSupported ErrorCode = 100002
	ErrorCodeNoRights     ErrorCode = 100003
	ErrorCodeBusy         ErrorCode = 100004

	ErrorCodeCSRF       ErrorCode = 125002
	ErrorCodeWrongToken ErrorCode = 125001
	// ErrorCodeWrongSession      ErrorCode = 125002
	ErrorCodeWrongSessionToken ErrorCode = 125003

	ErrorCodeUsernameWrong      ErrorCode = 108001
	ErrorCodePasswordWrong      ErrorCode = 108002
	ErrorCodeAlreadyLogin       ErrorCode = 108003
	ErrorCodeUsernamePwdWrong   ErrorCode = 108006
	ErrorCodeUsernamePwdOrerrun ErrorCode = 108007
	ErrorCodeUsernamePwdModify  ErrorCode = 115002
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
