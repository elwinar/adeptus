package adeptus

import (
	"reflect"
	"testing"
)

func Test_ParseUpgrade(t *testing.T) {
	cases := []struct {
		in  Line
		out RawUpgrade
		err bool
	}{
		{
			in:  Line{Text: ""},
			out: RawUpgrade{},
			err: true,
		},
		{
			in: Line{Text: "	"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in: Line{Text: " 	 "},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "fail"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "x fail"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "*"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "* [250]"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "* fail [250] fail"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "* fail [ 250]"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "* fail [abba250]"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "* fail 250]"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in:  Line{Text: "* fail [250"},
			out: RawUpgrade{},
			err: true,
		},
		{
			in: Line{Text: "* success"},
			out: RawUpgrade{
				mark: "*",
				name: "success",
			},
			err: false,
		},
		{
			in: Line{Text: "* success confirmed!"},
			out: RawUpgrade{
				mark: "*",
				name: "success confirmed!",
			},
			err: false,
		},
		{
			in: Line{Text: "* success 	(confirmed)"},
			out: RawUpgrade{
				mark: "*",
				name: "success (confirmed)",
			},
			err: false,
		},
		{
			in: Line{Text: "* [250] success"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: "* success [250]"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: "  * [250] success"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: "	* [250] success"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: "* success +4 [250]"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success +4",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: "* success	+4	[250]"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success +4",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: "* success: confirmed [250]"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success: confirmed",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: " * success - confirmed	[250]"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success - confirmed",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: "* success (confirmed) [250]"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success (confirmed)",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
		{
			in: Line{Text: " *	success	(confirmed) [250]"},
			out: RawUpgrade{
				mark:       "*",
				name:       "success (confirmed)",
				cost:       250,
				customCost: true,
			},
			err: false,
		},
	}

	for i, c := range cases {
		out, err := ParseUpgrade(c.in)
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
