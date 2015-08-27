package adeptus

import (
	"testing"
)

func Test_ParseSession(t *testing.T) {
	cases := []struct {
		in   string
		line int
		out  Session
		err  bool
	}{
		{
			in:   "",
			line: 1,
			out:  Session{line: 1},
			err:  true,
		},
		{
			in: "	",
			line: 1,
			out:  Session{line: 1},
			err:  true,
		},
		{
			in: " 	 ",
			line: 1,
			out:  Session{line: 1},
			err:  true,
		},
		{
			in: "fail",
			line: 1,
			out:  Session{line: 1},
			err:  true,
		},
		{
			in: "200 01 01",
			line: 1,
			out:  Session{line: 1},
			err:  true,
		},
		{
			in: "2001 04 28",
			line: 1,
			out:  Session{line: 1},
			err:  true,
		},
		{
			in: "2001_04_28",
			line: 1,
			out:  Session{line: 1},
			err:  true,
		},
		{
			in: "2001.04.28 [250",
			line: 1,
			out:  Session{Line: 1},
			err:  true,
		},
		{
			in: "2001.04.28 250]",
			line: 1,
			out:  Session{Line: 1},
			err:  true,
		},
		{
			in: "2001/04/28 success",
			line: 1,
			out:  Session{
				Line: 1,
				Date: time.Parse("2006/03/02", "2001/04/28"),
				Title: "success",
			},
			err:  false,
		},
		{
			in: "2001-04-28 success",
			line: 1,
			out:  Session{
				Line: 1,
				Date: time.Parse("2006/03/02", "2001/04/28"),
				Title: "success",
			},
			err:  false,
		},
		{
			in: "2001.04.28 success",
			line: 1,
			out:  Session{
				Line: 1,
				Date: time.Parse("2006/03/02", "2001/04/28"),
				Title: "success",
			},
			err:  false,
		},
		{
			in: "2001.04.28 [250]",
			line: 1,
			out:  Session{
				Line: 1,
				Date: time.Parse("2006/03/02", "2001/04/28"),
				Reward: 250,
			},
			err:  false,
		},
		{
			in: "2001.04.28 [250] success",
			line: 1,
			out:  Session{
				Line: 1,
				Date: time.Parse("2006/03/02", "2001/04/28"),
				Title: "success",
				Reward: 250,
			},
			err:  false,
		},
		{
			in: "2001.04.28 success [250]",
			line: 1,
			out:  Session{
				Line: 1,
				Date: time.Parse("2006/03/02", "2001/04/28"),
				Title: "success",
				Reward: 250,
			},
			err:  false,
		},
		{
			in: "	2001.04.28	success	[250]",
			line: 1,
			out:  Session{
				Line: 1,
				Date: time.Parse("2006/03/02", "2001/04/28"),
				Title: "success",
				Reward: 250,
			},
			err:  false,
		},
	}

	for i, c := range cases {
		out, err := ParseSession(c.line, c.in)
		if (err != nil) != c.err {
			t.Logf("Unexpected error on case %d:", i+1)
			t.Logf("	Having %s", err)
			t.Fail()
			continue
		}
		if out != c.out {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}
}
