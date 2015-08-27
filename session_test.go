package adeptus

import (
	"reflect"
	"testing"
	"time"
)

func Test_ParseSession(t *testing.T) {
	cases := []struct {
		in  string
		out Session
		err bool
	}{
		{
			in:  "",
			out: Session{},
			err: true,
		},
		{
			in: "	",
			out: Session{},
			err: true,
		},
		{
			in: " 	 ",
			out: Session{},
			err: true,
		},
		{
			in:  "fail",
			out: Session{},
			err: true,
		},
		{
			in:  "200 01 01",
			out: Session{},
			err: true,
		},
		{
			in:  "2001 04 28",
			out: Session{},
			err: true,
		},
		{
			in:  "2001_04_28",
			out: Session{},
			err: true,
		},
		{
			in:  "2001.04.28 [250",
			out: Session{},
			err: true,
		},
		{
			in:  "2001.04.28 250]",
			out: Session{},
			err: true,
		},
		{
			in:  "2001.04.28 fail [250] fail",
			out: Session{},
			err: true,
		},
		{
			in:  "2001.04.28 fail [ 250]",
			out: Session{},
			err: true,
		},
		{
			in:  "2001.04.28 fail [abba250]",
			out: Session{},
			err: true,
		},
		{
			in: "2001/04/28 success",
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: "2001-04-28 success",
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: "2001.04.28 success",
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: "2001.04.28 [250]",
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Reward: 250,
			},
			err: false,
		},
		{
			in: "2001.04.28 [250] success",
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
		{
			in: "2001.04.28 success [250]",
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
		{
			in: "	2001.04.28	success	[250]",
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
	}

	for i, c := range cases {
		out, err := ParseSession(Line{Text: c.in, Number: 1})
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
