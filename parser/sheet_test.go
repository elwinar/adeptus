package parser

import (
	"os"
	"reflect"
	"testing"
)

func Test_ParseSheet(t *testing.T) {
	cases := []struct {
		in  string
		out Sheet
		err bool
	}{
		{
			in:  "tests/empty-file.40k",
			out: Sheet{},
			err: true,
		},
		{
			in:  "tests/comments-only.40k",
			out: Sheet{},
			err: true,
		},
		{
			in:  "tests/wrong-header.40k",
			out: Sheet{},
			err: true,
		},
		{
			in:  "tests/wrong-session.40k",
			out: Sheet{},
			err: true,
		},
		{
			in:  "tests/no-session.40k",
			out: Sheet{},
			err: false,
		},
		{
			in:  "tests/success.40k",
			out: Sheet{},
			err: false,
		},
	}

	for i, c := range cases {
		in, err := os.Open(c.in)
		if err != nil {
			t.Fatalf("Unable to open file %s: %s", c.in, err)
		}

		out, err := ParseSheet(in)
		if (err != nil) != c.err {
			t.Logf("Unexpected error on case %d:", i+1)
			t.Logf("	Having %s", err)
			t.Fail()
			continue
		}

		if !reflect.DeepEqual(out, c.out) {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}

}
