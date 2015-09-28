package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseSession(t *testing.T) {
	cases := []struct {
		in    []string
		out   Session
		err   bool
		panic bool
	}{
		{
			in:    []string{},
			out:   Session{},
			err:   false,
			panic: true,
		},
		{
			in: []string{
				"",
			},
			out:   Session{},
			err:   false,
			panic: true,
		},
		{
			in: []string{
				"	",
			},
			out:   Session{},
			err:   false,
			panic: true,
		},
		{
			in: []string{
				" 	 ",
			},
			out:   Session{},
			err:   false,
			panic: true,
		},
		{
			in: []string{
				"fail",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"200 01 01",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001 04 28",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001_04_28",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 [250",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 250]",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 fail [250] fail",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 fail [ 250]",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 fail [abba250]",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001/04/28 [250] fail [250]",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001/04/28 success",
			},
			out: Session{
				Date:     time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:    "success",
				Reward:   nil,
				Upgrades: []Upgrade{},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"2001-04-28 success",
			},
			out: Session{
				Date:     time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:    "success",
				Reward:   nil,
				Upgrades: []Upgrade{},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 success",
			},
			out: Session{
				Date:     time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:    "success",
				Reward:   nil,
				Upgrades: []Upgrade{},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 [250]",
			},
			out: Session{
				Date:     time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:    "",
				Reward:   IntP(250),
				Upgrades: []Upgrade{},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 [250] success",
			},
			out: Session{
				Date:     time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:    "success",
				Reward:   IntP(250),
				Upgrades: []Upgrade{},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28 success [250]",
			},
			out: Session{
				Date:     time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:    "success",
				Reward:   IntP(250),
				Upgrades: []Upgrade{},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"	2001.04.28	success	[250]",
			},
			out: Session{
				Date:     time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:    "success",
				Reward:   IntP(250),
				Upgrades: []Upgrade{},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28	success	[250]",
				"",
			},
			out:   Session{},
			err:   false,
			panic: true,
		},
		{
			in: []string{
				"2001.04.28	success	[250]",
				"fsrfwsf",
			},
			out:   Session{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28	success	[250]",
				"	* WP +5",
				"	* WS +5",
				"	* BS +5",
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: IntP(250),
				Upgrades: []Upgrade{
					Upgrade{
						Mark: "*",
						Name: "WP +5",
						Cost: nil,
						Line: 2,
					},
					Upgrade{
						Mark: "*",
						Name: "WS +5",
						Cost: nil,
						Line: 3,
					},
					Upgrade{
						Mark: "*",
						Name: "BS +5",
						Cost: nil,
						Line: 4,
					},
				},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"2001.04.28	success	[250]",
				"",
			},
			out:   Session{},
			err:   false,
			panic: true,
		},
	}

	for i, c := range cases {
		in := []line{}
		for number, text := range c.in {
			in = append(in, newLine(text, number+1))
		}

		out, err, panic := func() (out Session, err error, panic bool) {
			defer func() {
				if e := recover(); e != nil {
					panic = true
				}
			}()

			out, err = parseSession(in)
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
