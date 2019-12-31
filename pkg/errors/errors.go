package errors

import (
	"fmt"
)

const (
	EMPTYKEYCODE   = 100
	EMPTYVALUECODE = 101

	EMPTYKEYMSG   = "Empty key"
	EMPTYVALUEMSG = "Empty value"
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
