package parser

import (
	"testing"
)

func Test_in(t *testing.T) {
	cases := []struct {
		in    string
		slice []string
		out   bool
	}{
		{
			in:    "a",
			slice: []string{},
			out:   false,
		},
		{
			in:    "a",
			slice: []string{"b", "c"},
			out:   false,
		},
		{
			in:    "a",
			slice: []string{"a", "b", "c"},
			out:   true,
		},
		{
			in:    "b",
			slice: []string{"a", "b", "c"},
			out:   true,
		},
		{
			in:    "c",
			slice: []string{"a", "b", "c"},
			out:   true,
		},
		{
			in:    "d",
			slice: []string{"a", "b", "c"},
			out:   false,
		},
	}

	for i, c := range cases {
		out := in(c.in, c.slice)
		if out != c.out {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("Expected %t", c.out)
			t.Logf("Having %t", out)
			t.Fail()
		}
	}
}
