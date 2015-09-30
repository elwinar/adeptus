package main

import "fmt"

// ParserError is an error encountered when parsing the sheet
type ParserError struct {
	Line int
	Code ErrorCode
}

// NewParserError build a new error from the line and error code
func NewParserError(line int, code ErrorCode) ParserError {
	return ParserError{
		Line: line,
		Code: code,
	}
}

// Error implements the Error interface
func (e ParserError) Error() string {
	msg, found := errorMsgs[e.Code]
	if !found {
		panic(fmt.Sprintf("undefined error message for code %d", e.Code))
	}
	return fmt.Sprintf("line %d: %s", e.Line, msg)
}
