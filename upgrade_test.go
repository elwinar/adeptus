package adeptus

import (
	"testing"
)

func Test_ParseUpgrade(t *testing.T) {
	cases := []struct {
		in   string
		out  RawUpgrade
		err  bool
	}{
		{
			in:   "",
			out:  RawUpgrade{},
			err:  true,
		},
		{
			in: "	",
			out:  RawUpgrade{},
			err:  true,
		},
		{
			in: " 	 ",
			out:  RawUpgrade{},
			err:  true,
		},
		{
			in: "fail",
			out:  RawUpgrade{},
			err:  true,
		},
		{
			in: "x fail",
			out:  RawUpgrade{},
			err:  true,
		},
		{
			in: "*",
			out:  RawUpgrade{},
			err:  true,
		},
		{
			in: "* [250]",
			out: RawUpgrade{},
			err: true,
		},
		{
			in: "* fail [250] fail",
			out: RawUpgrade{},
			err: true,
		},
		{
			in: "* fail [ 250]",
			out: RawUpgrade{},
			err: true,
		},
		{
			in: "* fail [abba250]",
			out: RawUpgrade{},
			err: true,
		},
		{
			in: "* fail 250]",
			out: RawUpgrade{},
			err: true,
		},
		{
			in: "* fail [250",
			out: RawUpgrade{},
			err: true,
		},
		{
			in: "* success",
			out: RawUpgrade{
				mark: "*",
				name: "success",
			},
			err: false,
		},
		{
			in: "* [250] success",
			out: RawUpgrade{
				mark: "*",
				name: "success",
			},
			err: false,
		},
		{
			in: "* success [250]",
			out: RawUpgrade{
				mark: "*",
				name: "success",
				cost: "250",
			},
			err: false,
		},
		{
			in: "  * [250] success",
			out: RawUpgrade{
				mark: "*",
				name: "success",
				cost: "250",
			},
			err: false,
		},
		{
			in: "	* [250] success",
			out: RawUpgrade{
				mark: "*",
				name: "success",
				cost: "250",
			},
			err: false,
		},
		{
			in: "* success +4 [250]",
			out: RawUpgrade{
				mark: "*",
				name: "success +4",
				cost: "250",
			},
			err: false,
		},
		{
			in: "* success	+4	[250]",
			out: RawUpgrade{
				mark: "*",
				name: "success +4",
				cost: "250",
			},
			err: false,
		},
		{
			in: "* success: confirmed [250]",
			out: RawUpgrade{
				mark: "*",
				name: "success: confirmed",
				cost: "250",
			},
			err: false,
		},
		{
			in: " * success - confirmed	[250]",
			out: RawUpgrade{
				mark: "*",
				name: "success - confirmed",
				cost: "250",
			},
			err: false,
		},
		{
			in: "* success (confirmed) [250]",
			out: RawUpgrade{
				mark: "*",
				name: "success (confirmed)",
				cost: "250",
			},
			err: false,
		},
		{
			in: " *	success	(confirmed) [250]",
			out: RawUpgrade{
				mark: "*",
				name: "success (confirmed)",
				cost: "250",
			},
			err: false,
		},
	}

	for i, c := range cases {
		out, err := ParseUpgrade(Line{Text: c.in, Number: 1})
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
