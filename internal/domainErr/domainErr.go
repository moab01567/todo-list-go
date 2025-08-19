package domainErr

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

type Code string

const (
	CodeNotFound     Code = "NOT_FOUND"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeInternal     Code = "INTERNAL_SERVER_ERROR"
)

type DomainError struct {
	Msg   string
	Code  Code
	Err   error
	Stack string
}

func New(msg string, err error, code Code) *DomainError {
	if msg == "" {
		msg = string(code)
	}
	return &DomainError{Msg: msg, Err: err, Code: code, Stack: string(debug.Stack())}
}
func (e *DomainError) PrintError() {
	fmt.Printf("%v\n%v\n%v\n", e.Stack, e.Err, e.Msg)
}

func (e *DomainError) Error() string {
	return e.Msg
}

func (e *DomainError) Unwrap() error {
	return e.Err
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
	case CodeNotFound:
		return http.StatusNotFound
	case CodeUnauthorized:
		return http.StatusUnauthorized
	default:
		domainErr.PrintError()
		return http.StatusInternalServerError
	}

}
