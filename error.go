package main

import (
	"fmt"
)

const (
	// UndefinedName is returned when character's name is not defined
	UndefinedName = 000

	// UndefinedUniverse is returned when the universe is not defined
	UndefinedUniverse = 001

	// UndefinedOrigin is returned when the origin is not defined
	UndefinedOrigin = 003

	// UndefinedBackground is returned when the background is not defined
	UndefinedBackground = 004

	// UndefinedRole is returned when the role is not defined
	UndefinedRole = 005

	// NotFoundUniverse is returned when the universe file cannot be opened
	NotFoundUniverse = 100

	// InvalidUniverse is returned when the universe is not a valid json file
	InvalidUniverse = 101
)

// ErrorCode holds the type of error return
type ErrorCode int

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
