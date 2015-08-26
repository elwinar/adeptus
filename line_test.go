package main

import(
	"reflect"
	"testing"
)

func Test_NewLine(t *testing.T) {
	
	cases := []struct{
		raw string,
		out *Line,
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
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasDate(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
	}{
		{
			line: "201/01/01",
			out: false,
		},
		{
			line: "2015 01 01",
			out: false,
		},
		{
			line: "date: 2015/01/01",
			out: false,
		},
		{
			line: "2015/01/01",
			out: true,
		},
		{
			line: "2015-01-01",
			out: true,
		},
		{
			line: "2015.01.01",
			out: true,
		},
		{
			line: "\s\s2015.01.01 I will survive 250xp",
			out: true,
		},
		{
			line: "\t\t2015.01.01 I will survive 250xp",
			out: true,
		},
	}
	
	for k, c := range cases {		
		out := c.line.HasDate()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetDate(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
		raw string
	}{
		{
			line: "201/01/01",
			out: "",
			raw: "201/01/01",
		},
		{
			line: "2015 01 01",
			out: "",
			raw: "2015 01 01",
		},
		{
			line: "date: 2015/01/01",
			out: "",
			raw: "date: 2015/01/01"
		},
		{
			line: "2015/01/01",
			out: "2015/01/01",
			raw: "",
		},
		{
			line: "2015-01-01",
			out: "2015/01/01",
			raw: "",
		},
		{
			line: "2015.01.01",
			out: "2015/01/01",
			raw: "",
		},
		{
			line: "\s\s2015.01.01",
			out: "2015/01/01",
			raw: "",
		},
		{
			line: "\t\t2015.01.01",
			out: "2015/01/01",
			raw: "",
		},
		{
			line: "\t\t2015.01.01 I will survive 250xp\s\s",
			out: "2015/01/01",
			raw: "\sI will survive 250xp\s\s",
		},
	}
	
	for k, c := range cases {		
		out := c.line.GetDate()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasXp(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
	}{
		{
			line: "",
			out: false,
		},
		{
			line: "250",
			out: false,
		},
		{
			line: "250xp",
			out: true,
		},
		{
			line: "(250xp)",
			out: true,
		},
		{
			line: "\t\t2015.01.01 I will survive 250xp\s\s",
			out: true,
		},
		{
			line: "\t\t2015.01.01 250xp\s\s",
			out: true,
		},
	}
	
	for k, c := range cases {		
		out := c.line.HasXp()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetXp(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
		raw string
	}{
		{
			line: "250",
			out: "",
			raw: "250",
		},
		{
			line: "250x",
			out: "",
			raw: "250x",
		},
		{
			line: "250xp",
			out: "250",
			raw: "",
		},
		{
			line: "(250xp)",
			out: "250",
			raw: "",
		},
		{
			line: "\t\t2015.01.01 I will survive 250xp\s\s",
			out: "250",
			raw: "\t\t2015.01.01 I will survive\s\s\s",
		},
		{
			line: "\t\t2015.01.01 250xp\s\s",
			out: "250",
			raw: "\t\t2015.01.01\s\s\s",
		},
		{
			line: "\t\t2015.01.01 (250xp)\s\s",
			out: "250",
			raw: "\t\t2015.01.01\s\s\s",
		},
	}
	
	for k, c := range cases {		
		out := c.line.GetXp()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasMark(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
	}{
		{
			line: "",
			out: false,
		},
		{
			line: "+",
			out: true,
		},
		{
			line: "-",
			out: true,
		},
		{
			line: "*",
			out: true,
		},
		{
			line: "\t- Dodge",
			out: true,
		},
		{
			line: "\t* Dodge",
			out: true,
		},
		{
			line: "\t+ WP +2",
			out: true,
		},
	}
	
	for k, c := range cases {		
		out := c.line.HasMark()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetMark(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
		raw string
	}{
		{
			line: "250",
			out: "",
			raw: "250",
		},
		{
			line: "+",
			out: "+",
			raw: "",
		},
		{
			line: "-",
			out: "-",
			raw: "",
		},
		{
			line: "*",
			out: "*",
			raw: "",
		},
		{
			line: "\t+ WP +5",
			out: "+",
			raw: "\t WP +5"",
		},
	}
	
	for k, c := range cases {		
		out := c.line.GetMark()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasValue(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
	}{
		{
			line: "",
			out: false,
		},
		{
			line: "250",
			out: false,
		},
		{
			line: "+",
			out: false,
		},
		{
			line: "-",
			out: false,
		},
		{
			line: "*",
			out: false,
		},
		{
			line: "\t + 25",
			out: false,
		},
		{
			line: "25",
			out: true,
		},
		{
			line: "\t+4",
			out: true,
		},
		{
			line: "\t4",
			out: true,
		},
		{
			line: "\t-4",
			out: true,
		},
	}
	
	for k, c := range cases {		
		out := c.line.HasValue()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetValue(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
		raw string
	}{
		{
			line: "250",
			out: "",
			raw: "250",
		},
		{
			line: "\t* WP 50",
			out: "50",
			raw: "\t* WP ",
		},
		{
			line: "WP + 50",
			out: "",
			raw: "WP + 50",
		},
		{
			line: "WP +50",
			out: "+50",
			raw: "WP ",
		},
		{
			line: "\t+ WP -50",
			out: "-50",
			raw: "\t+ WP ",
		},
		{
			line: "2015/12/08 Kill the King 350xp",
			out: "",
			raw: "2015/12/08 Kill the King 350xp",
		},
		{
			line: "\t * STR +5 300xp",
			out: "+5",
			raw: "\t * STR  300xp",
		},
	}
	
	for k, c := range cases {		
		out := c.line.GetValue()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_HasKey(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
	}{
		{
			line: "a",
			out: false,
		},
		{
			line: "a:",
			out: true,
		},
		{
			line: "a:b",
			out: true,
		},
		{
			line: "\ta:\sb",
			out: true,
		},
		{
			line: "\ta:\tb",
			out: true,
		},
	}
	
	for k, c := range cases {		
		out := c.line.HasKey()
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetKey(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
		raw string
	}{
		{
			line: "a",
			out: "",
			raw: "a",
		},
		{
			line: "a:b",
			out: "a",
			raw: "b",
		},
		{
			line: "\ta:\sb",
			out: "a",
			raw: "\sb",
		},
	}
	
	for k, c := range cases {		
		out := c.line.GetKey()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
		if raw != c.raw {
			t.Logf("Unexpected raw in case %d.", k)
			t.Logf("\tExpected %v.", c.raw)
			t.Logf("\tHaving %v.", raw)
			t.Fail()
			continue
		}
	}
}

func Test_Line_GetLabel(t *testing.T) {
	
	cases := []struct{
		line *Line
		out bool
		raw string
	}{
		{
			line: "",
			out: "",
		},
		{
			line: "a",
			out: "a",
		},
		{
			line: "\ta\s\n\r",
			out: "a",
		},
	}
	
	for k, c := range cases {		
		out := c.line.GetLabel()
		raw := c.line.raw
		if out != c.out {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
			continue
		}
	}
}