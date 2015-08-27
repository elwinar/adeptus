package adeptus

import (
	"testing"
)

func Test_ParseUpgrade(t *testing.T) {
	cases := []struct {
		in   string
		line int
		out  RawUpgrade
		err  bool
	}{
		{
			in:   "",
			line: 1,
			out:  RawUpgrade{line: 1},
			err:  true,
		},
		{
			in: "	",
			line: 1,
			out:  RawUpgrade{line: 1},
			err:  true,
		},
		{
			in: " 	 ",
			line: 1,
			out:  RawUpgrade{line: 1},
			err:  true,
		},
		{
			in:   "fail",
			line: 1,
			out:  RawUpgrade{line: 1},
			err:  true,
		},
		{
			in:   "x fail",
			line: 1,
			out:  RawUpgrade{line: 1},
			err:  true,
		},
		{
			in:   "* 250xp",
			line: 1,
			out: RawUpgrade{
				line: 1,
			},
			err: true,
		},
		{
			in:   "* success",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success",
			},
			err: false,
		},
		{
			in:   "* experience success",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "experience success",
			},
			err: false,
		},
		{
			in:   "* 250XP success",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success",
				cost: "250",
			},
			err: false,
		},
		{
			in:   "* 250xp success",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success",
				cost: "250",
			},
			err: false,
		},
		{
			in: "	* 250xp success",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success",
				cost: "250",
			},
			err: false,
		},
		{
			in:   "* success 250xp",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success",
				cost: "250",
			},
			err: false,
		},
		{
			in:   "* success (250xp)",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success",
				cost: "250",
			},
			err: false,
		},
		{
			in:   "* success +4 (250xp)",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success +4",
				cost: "250",
			},
			err: false,
		},
		{
			in: "* success	+4	(250xp)",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success +4",
				cost: "250",
			},
			err: false,
		},
		{
			in: "* success: confirmed	(250xp)",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success: confirmed",
				cost: "250",
			},
			err: false,
		},
		{
			in: " * success - confirmed	(250xp)",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success - confirmed",
				cost: "250",
			},
			err: false,
		},
		{
			in: "* success (confirmed)	(250xp)",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success (confirmed)",
				cost: "250",
			},
			err: false,
		},
		{
			in: " *	success	(confirmed)	(250xp)",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success (confirmed)",
				cost: "250",
			},
			err: false,
		},
	}

	for i, c := range cases {
		out, err := ParseUpgrade(c.line, c.in)
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
