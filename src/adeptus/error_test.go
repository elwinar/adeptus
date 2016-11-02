package main

import (
	"fmt"
	"testing"
)

func Test_NewError(t *testing.T) {
	err := NewError(UnitTest)

	if err.Code != UnitTest {
		t.Logf("invalid code: expected %d, got %d", UnitTest, err.Code)
		t.Fail()
	}
}

func Test_Error_Error(t *testing.T) {
	err := NewError(UnitTest, "test")

	if err.Error() != fmt.Sprintf(errorMsgs[UnitTest], "test") {
		t.Logf("invalid output: %s", err.Error())
		t.Fail()
	}
}

func Test_Error_PanicIfNotFound(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			return
		}
		t.Fail()
	}()

	err := NewError(ErrorCode(-1))
	_ = err.Error()
}
