package main

import (
	"fmt"
	"testing"
)

func Test_NewError(t *testing.T) {
	err := NewError(UnusedAptitude)

	if err.Code != UnusedAptitude {
		t.Logf("invalid code: expected %d, got %d", UnusedAptitude, err.Code)
		t.Fail()
	}
}

func Test_Error_Error(t *testing.T) {
	err := NewError(UnusedAptitude)

	if err.Error() != fmt.Sprintf("line 0: %s", errorMsgs[UnusedAptitude]) {
		t.Logf("invalid output: %s", err.Error())
		t.Fail()
	}
}
