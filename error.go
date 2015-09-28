package main

import (
	"fmt"
)

// Error is an error encountered when parsing the sheet
type Error struct {
	Code ErrorCode
	Text string
}

// NewError build a new error from the error code
func NewError(code ErrorCode, errs ...error) Error {

	var text string
	for _, e := range errs {
		text = fmt.Sprintf("%s: %s", text, e.Error())
	}
	return Error{
		Code: code,
		Text: text,
	}
}

// Error implements the Error interface
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Text)
}
