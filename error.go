package dynalo

import (
	"fmt"
)

type dynamicLogicError struct {
	code    string
	message string
}

const (
	jsonInvalidCode = "1"
	pathInvalidCode = "2"
)

func newError(code string, message string) error {
	return &dynamicLogicError{
		code:    code,
		message: message,
	}
}

func invalidPath(json string, path string) error {
	return newError(pathInvalidCode, fmt.Sprintf("can not query: %s from json: %s", path, json))
}

func invalidJson() error {
	return newError(jsonInvalidCode, "invalid json")
}

func (e *dynamicLogicError) Error() string {
	return e.message
}
