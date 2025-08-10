package domainErr

import (
	"errors"
	"net/http"
)

type Code string

const (
	CodeInvalid      Code = "INVALID"
	CodeNotFound     Code = "NOT_FOUND"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeInternal     Code = "INTERNAL"
)

type DomainError struct {
	Msg    string
	Public string
	Code   Code
	Err    error
}

func (e *DomainError) Error() string {
	return e.Msg
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

func New(public string, msg string, err error, code Code) error {
	return &DomainError{Msg: msg, Err: err, Code: code}
}

func AS(err error) (*DomainError, bool) {
	var domainErr *DomainError
	ok := errors.As(err, &domainErr)
	return domainErr, ok
}

func GetHttpStatus(err error) int {
	domainErr, ok := AS(err)
	if !ok {
		return http.StatusInternalServerError
	}

	switch domainErr.Code {
	case CodeInternal:
		return http.StatusInternalServerError
	case CodeNotFound:
		return http.StatusNotFound
	case CodeUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}

}
