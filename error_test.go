package main

import (
	"fmt"
	"testing"
)

func Test_NewError(t *testing.T) {
	err := NewError(0, InsuficientData)

	if err.Line != 0 {
		t.Logf("invalid line number: expected %d, got %d", 0, err.Line)
		t.Fail()
	}

	if err.Code != InsuficientData {
		t.Logf("invalid code: expected %d, got %d", InsuficientData, err.Code)
		t.Fail()
	}
}

func Test_Error_Error(t *testing.T) {
	err := NewError(0, InsuficientData)

	if err.Error() != fmt.Sprintf("line 0: %s", errorMsgs[InsuficientData]) {
		t.Logf("invalid output: %s", err.Error())
		t.Fail()
	}
}
