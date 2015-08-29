package parser

import (
	"reflect"
	"testing"
)

func Test_ParseUpgrade(t *testing.T) {
	cases := []struct {
		in  line
		out Upgrade
		err bool
	}{
		{
			in:  line{Text: ""},
			out: Upgrade{},
			err: true,
		},
		{
			in: line{Text: "	"},
			out: Upgrade{},
			err: true,
		},
		{
			in: line{Text: " 	 "},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "fail"},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "x fail"},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "*"},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "* [250]"},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "* fail [250] fail"},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "* fail [ 250]"},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "* fail [abba250]"},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "* fail 250]"},
			out: Upgrade{},
			err: true,
		},
		{
			in:  line{Text: "* fail [250"},
			out: Upgrade{},
			err: true,
		},
		{
			in: line{Text: "* success"},
			out: Upgrade{
				Mark: "*",
				Name: "success",
			},
			err: false,
		},
		{
			in: line{Text: "* success confirmed!"},
			out: Upgrade{
				Mark: "*",
				Name: "success confirmed!",
			},
			err: false,
		},
		{
			in: line{Text: "* success 	(confirmed)"},
			out: Upgrade{
				Mark: "*",
				Name: "success (confirmed)",
			},
			err: false,
		},
		{
			in: line{Text: "* [250] success"},
			out: Upgrade{
				Mark: "*",
				Name: "success",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: "* success [250]"},
			out: Upgrade{
				Mark: "*",
				Name: "success",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: "  * [250] success"},
			out: Upgrade{
				Mark: "*",
				Name: "success",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: "	* [250] success"},
			out: Upgrade{
				Mark: "*",
				Name: "success",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: "* success +4 [250]"},
			out: Upgrade{
				Mark: "*",
				Name: "success +4",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: "* success	+4	[250]"},
			out: Upgrade{
				Mark: "*",
				Name: "success +4",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: "* success: confirmed [250]"},
			out: Upgrade{
				Mark: "*",
				Name: "success: confirmed",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: " * success - confirmed	[250]"},
			out: Upgrade{
				Mark: "*",
				Name: "success - confirmed",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: "* success (confirmed) [250]"},
			out: Upgrade{
				Mark: "*",
				Name: "success (confirmed)",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: line{Text: " *	success	(confirmed) [250]"},
			out: Upgrade{
				Mark: "*",
				Name: "success (confirmed)",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in:  line{Text: " * fail [-250]"},
			out: Upgrade{},
			err: true,
		},
	}

	for i, c := range cases {
		out, err := parseUpgrade(c.in)

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
