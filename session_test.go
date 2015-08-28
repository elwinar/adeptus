package adeptus

import (
	"reflect"
	"testing"
	"time"
)

func Test_ParseSession(t *testing.T) {
	cases := []struct {
		in  []Line
		out Session
		err bool
	}{
		{
			in:  []Line{},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:""}},
			out: Session{},
			err: true,
		},
		{
			in: []Line{Line{Text:"	"}},
			out: Session{},
			err: true,
		},
		{
			in: []Line{Line{Text:" 	 "}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"fail"}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"200 01 01"}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001 04 28"}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001_04_28"}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 [250"}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 250]"}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 fail [250] fail"}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 fail [ 250]"}},
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 fail [abba250]"}},
			out: Session{},
			err: true,
		},
		{
			in: []Line{Line{Text:"2001/04/28 success"}},
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001-04-28 success"}},
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001.04.28 success"}},
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001.04.28 [250]"}},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Reward: 250,
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001.04.28 [250] success"}},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001.04.28 success [250]"}},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"	2001.04.28	success	[250]"}},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
		{
			in: []Line{
				Line{Text:"2001.04.28	success	[250]"}},
				Line{Text:"	* WP +5"}},
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
		{
			in: []Line{
				Line{Text:"2001.04.28	success	[250]"}},
				Line{Text:"	* WP +5"}},
				Line{Text:"	x WP +5"}}, // provokes an error: incorrect mark
			},
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: true,
		},
	}

	for i, c := range cases {
		out, err := ParseSession(c.in)
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
