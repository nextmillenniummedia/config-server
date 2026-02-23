package errors

import (
	errors_std "errors"
	"fmt"
)

type HttpError struct {
	Code    int
	Message string
}

var CommandNotFoundError = errors_std.New("Command not found")
var BadRequestError = errors_std.New("Bad request")

func CommandNotFound(commandName string) error {
	return fmt.Errorf("%w: %s", CommandNotFoundError, commandName)
}
