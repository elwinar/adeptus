package main

import (
	"reflect"
	"testing"
)

func Test_parseUpgrade(t *testing.T) {
	cases := []struct {
		in    string
		out   Upgrade
		err   bool
		panic bool
	}{
		{
			in:    "",
			out:   Upgrade{},
			err:   false,
			panic: true,
		},
		{
			in: "	",
			out:   Upgrade{},
			err:   false,
			panic: true,
		},
		{
			in: " 	 ",
			out:   Upgrade{},
			err:   false,
			panic: true,
		},
		{
			in:  "fail",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "x fail",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "*",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "* [250]",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "* fail [250] fail",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "* fail [ 250]",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "* fail [abba250]",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "* fail 250]",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "* fail [250",
			out: Upgrade{},
			err: true,
		},
		{
			in:  "* [250] fail [250]",
			out: Upgrade{},
			err: true,
		},
		{
			in: "* success",
			out: Upgrade{
				Mark: "*",
				Name: "success",
			},
			err: false,
		},
		{
			in: "* success confirmed!",
			out: Upgrade{
				Mark: "*",
				Name: "success confirmed!",
			},
			err: false,
		},
		{
			in: "* success 	(confirmed)",
			out: Upgrade{
				Mark: "*",
				Name: "success (confirmed)",
			},
			err: false,
		},
		{
			in: "* [250] success",
			out: Upgrade{
				Mark: "*",
				Name: "success",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: "* success [250]",
			out: Upgrade{
				Mark: "*",
				Name: "success",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: "  * [250] success",
			out: Upgrade{
				Mark: "*",
				Name: "success",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: "	* [250] success",
			out: Upgrade{
				Mark: "*",
				Name: "success",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: "* success +4 [250]",
			out: Upgrade{
				Mark: "*",
				Name: "success +4",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: "* success	+4	[250]",
			out: Upgrade{
				Mark: "*",
				Name: "success +4",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: "* success: confirmed [250]",
			out: Upgrade{
				Mark: "*",
				Name: "success: confirmed",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: " * success - confirmed	[250]",
			out: Upgrade{
				Mark: "*",
				Name: "success - confirmed",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: "* success (confirmed) [250]",
			out: Upgrade{
				Mark: "*",
				Name: "success (confirmed)",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in: " *	success	(confirmed) [250]",
			out: Upgrade{
				Mark: "*",
				Name: "success (confirmed)",
				Cost: IntP(250),
			},
			err: false,
		},
		{
			in:  " * fail [-250]",
			out: Upgrade{},
			err: true,
		},
	}

	for i, c := range cases {
		out, err, panic := func() (out Upgrade, err error, panic bool) {
			defer func() {
				if e := recover(); e != nil {
					panic = true
				}
			}()

			out, err = parseUpgrade(newLine(c.in, 0))
			return
		}()

		if (err != nil) != c.err {
			if err == nil {
				t.Logf("Expected error on case %d", i+1)
			} else {
				t.Logf("Unexpected error on case %d: %s", i+1, err)
			}
			t.Fail()
		}

		if !reflect.DeepEqual(out, c.out) {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}

		if panic != c.panic {
			if panic {
				t.Logf("Unexpected panic on case %d", i+1)
			} else {
				t.Logf("Should panic on case %d", i+1)
			}
			t.Fail()
		}
	}
}
