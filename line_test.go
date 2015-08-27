package adeptus

import (
	"testing"
)

func Test_Line_IsEmpty(t *testing.T) {

	cases := []struct {
		in  Line
		out bool
	}{
		{
			in:  Line{Text: "Not empty"},
			out: false,
		},
		{
			in:  Line{Text: " "},
			out: false,
		},
		{
			in: Line{Text: "	"},
			out: false,
		},
		{
			in:  Line{Text: ""},
			out: true,
		},
	}

	for i, c := range cases {
		out := c.in.IsEmpty()
		if out != c.out {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}
}

func Test_Line_IsComment(t *testing.T) {

	cases := []struct {
		in  Line
		out bool
	}{
		{
			in:  Line{Text: "Not a comment"},
			out: false,
		},
		{
			in:  Line{Text: "* Not a comment"},
			out: false,
		},
		{
			in:  Line{Text: "+ Not a comment"},
			out: false,
		},
		{
			in:  Line{Text: "/ / Not a comment"},
			out: false,
		},
		{
			in:  Line{Text: "// A comment"},
			out: true,
		},
		{
			in:  Line{Text: "# A comment"},
			out: true,
		},
		{
			in:  Line{Text: "; A comment"},
			out: true,
		},
	}

	for i, c := range cases {
		out := c.in.IsComment()
		if out != c.out {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}
}
