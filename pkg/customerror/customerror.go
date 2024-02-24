package customerror

import (
	"errors"
	"fmt"
)

type Error struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
}

func New(code ErrCode, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("[Error %d]: %s", e.Code, e.Message)
}

func FromError(err error) *Error {
	var customErr *Error
	ok := errors.As(err, &customErr)
	if !ok {
		return &Error{
			Code:    UnknownErrorCode,
			Message: err.Error(),
		}
	}

	fmt.Println(customErr)
	return err.(*Error)
}

func Wrap(f string, err error) error {
	return fmt.Errorf(f, err)
}

func Unwrap(err error) error {
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	return err
}
