package appError

import (
	"fmt"
	"runtime/debug"
)

type AppError struct {
	Message string
	Stack   string
	Err     error
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s\n %s\n %v\n", e.Message, e.Stack, e.Err)
}
func New(message string, err error) error {
	return AppError{Message: message, Err: err, Stack: string(debug.Stack())}
}
