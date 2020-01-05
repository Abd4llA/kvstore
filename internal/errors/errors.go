package errors

import (
	"fmt"
)

const (
	EMPTY_KEY_CODE   = 100
	EMPTY_VALUE_CODE = 101

	EMPTY_KEY_MSG   = "Empty key"
	EMPTY_VALUE_MSG = "Empty value"
)

type KVSError struct {
	text string
	code uint
}

func (error KVSError) Error() string {
	return fmt.Sprintf("Error Code: %v\n Error Message: %v", error.code, error.text)
}

func New(text string, code uint) KVSError {
	return KVSError{
		text: text,
		code: code,
	}
}
