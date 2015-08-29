package parser

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseSession(t *testing.T) {
	cases := []struct {
		in  []string
		out Session
		err bool
	}{
		{
			in:  []string{},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"	",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				" 	 ",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"fail",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"200 01 01",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"2001 04 28",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"2001_04_28",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"2001.04.28 [250",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"2001.04.28 250]",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"2001.04.28 fail [250] fail",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"2001.04.28 fail [ 250]",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"2001.04.28 fail [abba250]",
			},
			out: Session{},
			err: true,
		},
		{
			in: []string{
				"2001/04/28 success",
			},
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []string{
				"2001-04-28 success",
			},
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []string{
				"2001.04.28 success",
			},
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []string{
				"2001.04.28 [250]",
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Reward: IntP(250),
			},
			err: false,
		},
		{
			in: []string{
				"2001.04.28 [250] success",
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: IntP(250),
			},
			err: false,
		},
		{
			in: []string{
				"2001.04.28 success [250]",
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: IntP(250),
			},
			err: false,
		},
		{
			in: []string{
				"	2001.04.28	success	[250]",
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: IntP(250),
			},
			err: false,
		},
		{
			in: []string{
				"2001.04.28	success	[250]",
				"",
				"",
				"",
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: IntP(250),
				Upgrades: []Upgrade{
					Upgrade{},
					Upgrade{},
					Upgrade{},
				},
			},
			err: false,
		},
		{
			in: []string{
				"2001.04.28	success	[250]",
				"",
				"",
				"",
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: IntP(250),
				Upgrades: []Upgrade{
					Upgrade{},
					Upgrade{},
					Upgrade{},
				},
			},
			err: false,
		},
		{
			in: []string{
				"2001.04.28	success	[250]",
				"",
			},
			out: Session{},
			err: true,
		},
	}

	for i, c := range cases {
		in := []line{}
		for number, text := range c.in {
			in = append(in, newLine(text, number))
		}

		out, err := parseSession(in)
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
