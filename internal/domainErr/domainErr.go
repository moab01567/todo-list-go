package domainErr

import (
	"errors"
)

type Code string

const (
	CodeInvalid      Code = "INVALID"
	CodeNotFound     Code = "NOT_FOUND"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeInternal     Code = "INTERNAL"
)

type DomainError struct {
	Msg  string
	Code Code
	Err  error
}

func (e *DomainError) Error() string {
	return e.Msg
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

func New(msg string, err error, code Code) *DomainError {
	return &DomainError{Msg: msg, Err: err, Code: code}
}

func Wrap(msg string, err error, code Code) error {
	return &DomainError{Msg: msg, Err: err, Code: code}
}

func AS(err error) (*DomainError, bool) {
	var domainErr *DomainError
	ok := errors.As(err, &domainErr)
	return domainErr, ok
}
