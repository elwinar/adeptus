package main

import (
	"reflect"
	"testing"
)

func Test_NewMeta(t *testing.T) {
	cases := []struct {
		in  string
		out Meta
		err bool
	}{
		{
			in:  "",
			out: Meta{},
			err: false,
		},
		{
			in: "something",
			out: Meta{
				Label: "something",
			},
			err: false,
		},
	}

	for i, c := range cases {

		out, err := NewMeta(newLine(c.in, 0))

		if (err != nil) != c.err {
			if err == nil {
				t.Logf("Expected error on case %d", i+1)
			} else {
				t.Logf("Unexpected error on case %d: %s", i+1, err)
			}
			t.Fail()
		}

		if !reflect.DeepEqual(out, c.out) {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}
}
