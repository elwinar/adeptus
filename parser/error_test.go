package parser

import (
	"testing"
)

func Test_NewError(t *testing.T) {
	err := NewError(0, EmptySheet)
	
	if err.Line != 0 {
		t.Logf("invalid line number: expected %d, got %d", 0, err.Line)
		t.Fail()
	}
	
	if err.Code != EmptySheet {
		t.Logf("invalid code: expected %d, got %d", EmptySheet, err.Code)
		t.Fail()
	}
}

func Test_Error_Error(t *testing.T) {
	err := NewError(0, EmptySheet)
	
	if err.Error() != "line 0: 0" {
		t.Logf("invalid output: %s", err.Error())
		t.Fail()
	}
}
