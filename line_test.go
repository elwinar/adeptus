package main

import (
	"reflect"
	"testing"
)

func Test_NewLine(t *testing.T) {

	cases := []struct {
		raw string
		out *Line
	}{
		{
			raw: "",
			out: &Line{
				raw: "",
			},
		},
		{
			raw: "test",
			out: &Line{
				raw: "test",
			},
		},
		{
			raw: "The king must die",
			out: &Line{
				raw: "The king must die",
			},
		},
	}

	for k, c := range cases {
		out := NewLine(c.raw)
		if !reflect.DeepEqual(out, c.out) {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasDate(t *testing.T) {

	cases := []struct {
		line *Line
		out  bool
	}{
		{
			line: NewLine("201/01/01"),
			out:  false,
		},
		{
			line: NewLine("2015 01 01"),
			out:  false,
		},
		{
			line: NewLine("date: 2015/01/01"),
			out:  false,
		},
		{
			line: NewLine("2015/01/01"),
			out:  true,
		},
		{
			line: NewLine("2015-01-01"),
			out:  true,
		},
		{
			line: NewLine("2015.01.01"),
			out:  true,
		},
		{
			line: NewLine("  2015.01.01 I will survive 250xp"),
			out:  true,
		},
		{
			line: NewLine("\t\t2015.01.01 I will survive 250xp"),
			out:  true,
		},
	}

	for k, c := range cases {
		out := c.line.HasDate()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetDate(t *testing.T) {

	cases := []struct {
		line *Line
		out  string
		raw  string
	}{
		{
			line: NewLine("201/01/01"),
			out:  "",
			raw:  "201/01/01",
		},
		{
			line: NewLine("2015 01 01"),
			out:  "",
			raw:  "2015 01 01",
		},
		{
			line: NewLine("date: 2015/01/01"),
			out:  "",
			raw:  "date: 2015/01/01",
		},
		{
			line: NewLine("2015/01/01"),
			out:  "2015/01/01",
			raw:  "",
		},
		{
			line: NewLine("2015-01-01"),
			out:  "2015/01/01",
			raw:  "",
		},
		{
			line: NewLine("2015.01.01"),
			out:  "2015/01/01",
			raw:  "",
		},
		{
			line: NewLine("  2015.01.01"),
			out:  "2015/01/01",
			raw:  "",
		},
		{
			line: NewLine("\t\t2015.01.01"),
			out:  "2015/01/01",
			raw:  "",
		},
		{
			line: NewLine("\t\t2015.01.01 I will survive 250xp  "),
			out:  "2015/01/01",
			raw:  " I will survive 250xp  ",
		},
	}

	for k, c := range cases {
		out := c.line.GetDate()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k+1)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasXp(t *testing.T) {

	cases := []struct {
		line *Line
		out  bool
	}{
		{
			line: NewLine(""),
			out:  false,
		},
		{
			line: NewLine("250"),
			out:  false,
		},
		{
			line: NewLine("250xp"),
			out:  true,
		},
		{
			line: NewLine("(250xp)"),
			out:  true,
		},
		{
			line: NewLine("\t\t2015.01.01 I will survive 250xp  "),
			out:  true,
		},
		{
			line: NewLine("\t\t2015.01.01 250xp  "),
			out:  true,
		},
	}

	for k, c := range cases {
		out := c.line.HasXp()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetXp(t *testing.T) {

	cases := []struct {
		line *Line
		out  string
		raw  string
	}{
		{
			line: NewLine("250"),
			out:  "",
			raw:  "250",
		},
		{
			line: NewLine("250x"),
			out:  "",
			raw:  "250x",
		},
		{
			line: NewLine("250xp"),
			out:  "250",
			raw:  "",
		},
		{
			line: NewLine("(250xp)"),
			out:  "250",
			raw:  "",
		},
		{
			line: NewLine("\t\t2015.01.01 I will survive 250xp  "),
			out:  "250",
			raw:  "\t\t2015.01.01 I will survive   ",
		},
		{
			line: NewLine("\t\t2015.01.01 250xp  "),
			out:  "250",
			raw:  "\t\t2015.01.01   ",
		},
		{
			line: NewLine("\t\t2015.01.01 (250xp)  "),
			out:  "250",
			raw:  "\t\t2015.01.01   ",
		},
	}

	for k, c := range cases {
		out := c.line.GetXp()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k+1)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasMark(t *testing.T) {

	cases := []struct {
		line *Line
		out  bool
	}{
		{
			line: NewLine(""),
			out:  false,
		},
		{
			line: NewLine("+"),
			out:  true,
		},
		{
			line: NewLine("-"),
			out:  true,
		},
		{
			line: NewLine("*"),
			out:  true,
		},
		{
			line: NewLine("\t- Dodge"),
			out:  true,
		},
		{
			line: NewLine("\t* Dodge"),
			out:  true,
		},
		{
			line: NewLine("\t+ WP +2"),
			out:  true,
		},
	}

	for k, c := range cases {
		out := c.line.HasMark()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetMark(t *testing.T) {

	cases := []struct {
		line *Line
		out  string
		raw  string
	}{
		{
			line: NewLine("250"),
			out:  "",
			raw:  "250",
		},
		{
			line: NewLine("+"),
			out:  "+",
			raw:  "",
		},
		{
			line: NewLine("-"),
			out:  "-",
			raw:  "",
		},
		{
			line: NewLine("*"),
			out:  "*",
			raw:  "",
		},
		{
			line: NewLine("\t+ WP +5 (250xp)"),
			out:  "+",
			raw:  "\t WP +5 (250xp)",
		},
	}

	for k, c := range cases {
		out := c.line.GetMark()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k+1)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasValue(t *testing.T) {

	cases := []struct {
		line *Line
		out  bool
	}{
		{
			line: NewLine("\t* WP 250"),
			out:  false,
		},
		{
			line: NewLine(""),
			out:  false,
		},
		{
			line: NewLine("250"),
			out:  false,
		},
		{
			line: NewLine("+"),
			out:  false,
		},
		{
			line: NewLine("-"),
			out:  false,
		},
		{
			line: NewLine("*"),
			out:  false,
		},
		{
			line: NewLine("\t + 25"),
			out:  true,
		},
		{
			line: NewLine("25"),
			out:  true,
		},
		{
			line: NewLine("\t+4"),
			out:  true,
		},
		{
			line: NewLine("\t4"),
			out:  true,
		},
		{
			line: NewLine("\t-4"),
			out:  true,
		},
		{
			line: NewLine("\t+ WP +5 (250xp)"),
			out:  true,
		},
	}

	for k, c := range cases {
		out := c.line.HasValue()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetValue(t *testing.T) {

	cases := []struct {
		line *Line
		out  string
		raw  string
	}{
		{
			line: NewLine("\t* WP 250"),
			out:  "",
			raw:  "\t* WP 250",
		},
		{
			line: NewLine("250"),
			out:  "",
			raw:  "250",
		},
		{
			line: NewLine("\t* WP 50"),
			out:  "50",
			raw:  "\t* WP ",
		},
		{
			line: NewLine("WP + 50"),
			out:  "+50",
			raw:  "WP ",
		},
		{
			line: NewLine("WP +        50"),
			out:  "+50",
			raw:  "WP ",
		},
		{
			line: NewLine("WP +50"),
			out:  "+50",
			raw:  "WP ",
		},
		{
			line: NewLine("\t+ WP -50"),
			out:  "-50",
			raw:  "\t+ WP ",
		},
		{
			line: NewLine("2015/12/08 Kill the King 350xp"),
			out:  "",
			raw:  "2015/12/08 Kill the King 350xp",
		},
		{
			line: NewLine("\t * STR +5 300xp"),
			out:  "+5",
			raw:  "\t * STR  300xp",
		},
	}

	for k, c := range cases {
		out := c.line.GetValue()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k+1)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasKey(t *testing.T) {

	cases := []struct {
		line *Line
		out  bool
	}{
		{
			line: NewLine("a"),
			out:  false,
		},
		{
			line: NewLine("a:"),
			out:  true,
		},
		{
			line: NewLine("a:b"),
			out:  true,
		},
		{
			line: NewLine("\ta: b"),
			out:  true,
		},
		{
			line: NewLine("\ta:\tb"),
			out:  true,
		},
	}

	for k, c := range cases {
		out := c.line.HasKey()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetKey(t *testing.T) {

	cases := []struct {
		line *Line
		out  string
		raw  string
	}{
		{
			line: NewLine("a"),
			out:  "",
			raw:  "a",
		},
		{
			line: NewLine("a:b"),
			out:  "a",
			raw:  "b",
		},
		{
			line: NewLine("\ta: b"),
			out:  "a",
			raw:  " b",
		},
		{
			line: NewLine("A:b"),
			out:  "a",
			raw:  "b",
		},
	}

	for k, c := range cases {
		out := c.line.GetKey()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k+1)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetLabel(t *testing.T) {

	cases := []struct {
		line *Line
		out  string
	}{
		{
			line: NewLine(""),
			out:  "",
		},
		{
			line: NewLine("a"),
			out:  "a",
		},
		{
			line: NewLine("\ta \n\r"),
			out:  "a",
		},
	}

	for k, c := range cases {
		out := c.line.GetLabel()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k+1)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}
