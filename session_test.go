package adeptus

import (
	"reflect"
	"testing"
	"time"
)

func successUpgradeParser(_ Line) (Upgrade, error) {
	return Upgrade{}, nil
}

func failUpgradeParser(_ Line) (Upgrade, error) {
	return Upgrade{}, errors.New("fail")
}

func Test_parseSession(t *testing.T) {
	cases := []struct {
		in  []Line
		parser upgradeParser
		out Session
		err bool
	}{
		{
			in:  []Line{},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:""}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in: []Line{Line{Text:"	"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in: []Line{Line{Text:" 	 "}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"fail"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"200 01 01"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001 04 28"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001_04_28"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 [250"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 250]"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 fail [250] fail"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 fail [ 250]"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in:  []Line{Line{Text:"2001.04.28 fail [abba250]"}},
			parser: successUpgradeParser,
			out: Session{},
			err: true,
		},
		{
			in: []Line{Line{Text:"2001/04/28 success"}},
			parser: successUpgradeParser,
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001-04-28 success"}},
			parser: successUpgradeParser,
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001.04.28 success"}},
			parser: successUpgradeParser,
			out: Session{
				Date:  time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title: "success",
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001.04.28 [250]"}},
			parser: successUpgradeParser,
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Reward: 250,
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001.04.28 [250] success"}},
			parser: successUpgradeParser,
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"2001.04.28 success [250]"}},
			parser: successUpgradeParser,
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
			},
			err: false,
		},
		{
			in: []Line{Line{Text:"	2001.04.28	success	[250]"}},
			parser: successUpgradeParser,
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
					Line{},
					Line{},
					Line{},
				},
			},
			parser: successUpgradeParser,
			out: Session{
				Date:   time.Date(2001, time.April, 28, 0, 0, 0, 0, time.UTC),
				Title:  "success",
				Reward: 250,
				Upgrades: []Upgrade{
					Upgrade{},
					Upgrade{},
					Upgrade{},
				},
			},
			err: false,
		},
		{
			in: []Line{
					Line{Text:"2001.04.28	success	[250]"},
					Line{},
				}
			},
			parser: failUpgradeParser,
			out: Session{},
			err: true,
		},
	}

	for i, c := range cases {
		out, err := parseSession(c.in)
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
