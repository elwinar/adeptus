package main

import (
	"testing"
)

func Test_newLine(t *testing.T) {
	l := newLine("raw", 0)

	if l.Number != 0 {
		t.Logf("invalid line number: expected %d, got %d", 0, l.Number)
		t.Fail()
	}

	if l.Text != "raw" {
		t.Logf("invalid text: expected %s, got %s", "raw", l.Text)
		t.Fail()
	}
}

func Test_line_IsEmpty(t *testing.T) {

	cases := []struct {
		in  string
		out bool
	}{
		{
			in:  "Not empty",
			out: false,
		},
		{
			in:  " ",
			out: true,
		},
		{
			in: "	",
			out: true,
		},
		{
			in:  "",
			out: true,
		},
	}

	for i, c := range cases {
		out := newLine(c.in, 0).IsEmpty()

		if out != c.out {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}
}

func Test_line_Instruction(t *testing.T) {

	cases := []struct {
		in  string
		out string
	}{
		{
			in:  "Not a comment",
			out: "Not a comment",
		},
		{
			in:  "* Not a comment",
			out: "* Not a comment",
		},
		{
			in:  "+ Not a comment",
			out: "+ Not a comment",
		},
		{
			in:  "/ / Not a comment",
			out: "/ / Not a comment",
		},
		{
			in:  "// A comment",
			out: "",
		},
		{
			in:  "Comment before// Comment after",
			out: "Comment before",
		},
		{
			in:  "# A comment",
			out: "",
		},
		{
			in:  "Comment before# Comment between // Comment after",
			out: "Comment before",
		},
		{
			in:  "Comment before// Comment between # Comment after",
			out: "Comment before",
		},
	}

	for i, c := range cases {
		out := newLine(c.in, 0).Instruction()

		if out != c.out {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}
}
