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
	err := NewError(UnusedAptitude, "test")

	if err.Error() != fmt.Sprintf(errorMsgs[UnusedAptitude], "test") {
		t.Logf("invalid output: %s", err.Error())
		t.Fail()
	}
}
