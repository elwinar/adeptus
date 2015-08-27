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
			in: "fail",
			line: 1,
			out:  RawUpgrade{line: 1},
			err:  true,
		},
		{
			in: "x fail",
			line: 1,
			out:  RawUpgrade{line: 1},
			err:  true,
		},
		{
			in: "*",
			line: 1,
			out:  RawUpgrade{line: 1},
			err:  true,
		},
		{
			in: "* [250]",
			line: 1,
			out: RawUpgrade{
				line: 1,
			},
			err: true,
		},
		{
			in: "* fail [250] fail",
			line: 1,
			out: RawUpgrade{
				line: 1,
			},
			err: true,
		},
		{
			in: "* fail [ 250]",
			line: 1,
			out: RawUpgrade{
				line: 1,
			},
			err: true,
		},
		{
			in: "* fail [abba250]",
			line: 1,
			out: RawUpgrade{
				line: 1,
			},
			err: true,
		},
		{
			in: "* fail 250]",
			line: 1,
			out: RawUpgrade{
				line: 1,
			},
			err: true,
		},
		{
			in: "* fail [250",
			line: 1,
			out: RawUpgrade{
				line: 1,
			},
			err: true,
		},
		{
			in: "* success",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success",
			},
			err: false,
		},
		{
			in: "* [250] success",
			line: 1,
			out: RawUpgrade{
				line: 1,
				mark: "*",
				name: "success",
			},
			err: false,
		},
		{
			in: "* success [250]",
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
			in: "  * [250] success",
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
			in: "	* [250] success",
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
			in: "* success +4 [250]",
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
			in: "* success	+4	[250]",
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
			in: "* success: confirmed [250]",
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
			in: " * success - confirmed	[250]",
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
			in: "* success (confirmed) [250]",
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
			in: " *	success	(confirmed) [250]",
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
