package commons

import (
	"fmt"
)

type ErrorDescription string

const (
	ErrorDescriptionInternalServerError ErrorDescription = "INTERNAL_SERVER_ERROR"
	ErrorDescriptionNotFound            ErrorDescription = "NOT_FOUND_ERROR"
	ErrorDescriptionBadUserInput        ErrorDescription = "BAD_USER_INPUT"
)

type ErrorCode string

const (
	ErrorCodeInternalServerError ErrorCode = "500"
	ErrorCodeNotFound            ErrorCode = "404"
	ErrorCodeBadUserInput        ErrorCode = "400"
)

type Error struct {
	Code        ErrorCode
	Description ErrorDescription
	Err         error
}

func NewError(code ErrorCode, description ErrorDescription, err error) *Error {
	return &Error{
		Code:        code,
		Description: description,
		Err:         err,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: %s", e.Err)
}
