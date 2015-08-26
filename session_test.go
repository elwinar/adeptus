package main

import(
	"testing"
)

func Test_Session_addLabel(t *testing.T) {
	
	cases := []struct{
		raw string,
		err bool,
		out Session,
	}{
		// errors
		{
			// nil input
			raw: "",
			err: true,
			out: Session{},
		},
		{
			// invalid date
			raw: "201/01/01\n",
			err: true,
			out: Session{},
		},
		{
			// xp without label or date
			raw: "250xp\n",
			err: true,
			out: Session{},
		},
		// date formats
		{
			raw: "01/01/2015\n",
			err: false,
			out: Session{
				date: "01/01/2015"
			},
		},
		{
			raw: "01-01-2015\n",
			err: false,
			out: Session{
				date: "01/01/2015"
			},
		},
		{
			raw: "01_01_2015\n",
			err: false,
			out: Session{
				date: "01/01/2015"
			},
		},
		{
			raw: "01.01.2015\n",
			err: false,
			out: Session{
				date: "01/01/2015"
			},
		},
		{
			raw: "01012015\n",
			err: false,
			out: Session{
				date: "01/01/2015"
			},
		},
		// raw analysis
		{
			// <date>( )<label>( )<xp>
			raw: "01/01/2015 The day Eddard died. 250xp\n",
			err: false,
			out: Session{
				date: "01/01/2015",
				label: "The day Eddard died.",
				xp: 250,
			},
		},
		{
			// <date>( )*<label>( )*<xp>
			raw: "01/01/2015    The day Eddard died.   250xp\n",
			err: false,
			out: Session{
				date: "01/01/2015",
				label: "The day Eddard died.",
				xp: 250,
			},
		},
		{
			// <date>(\t)<label>(\t)<xp>
			raw: "01/01/2015\tThe day Eddard died.\t250xp\n",
			err: false,
			out: Session{
				date: "01/01/2015",
				label: "The day Eddard died.",
				xp: 250,
			},
		},
		{
			// <date>(\t)*<label>(\t)*<xp>
			raw: "01/01/2015\t\tThe day Eddard died.\t\t\t250xp\n",
			err: false,
			out: Session{
				date: "01/01/2015",
				label: "The day Eddard died.",
				xp: 250,
			},
		},
		{
			// <date>
			raw: "01/01/2015\n",
			err: false,
			out: Session{
				date: "01/01/2015",
			},
		},
		{
			// <date> <label>
			raw: "01/01/2015 The day Eddard died.\n",
			err: false,
			out: Session{
				date: "01/01/2015",
				label: "The day Eddard died.",
			},
		},
		{
			// <date> <xp>
			raw: "01/01/2015 250xp\n",
			err: false,
			out: Session{
				date: "01/01/2015",
				xp: 250,
			},
		},
		{
			// <date> <label> <xp>
			raw: "01/01/2015 The day Eddard died. 250xp\n",
			err: false,
			out: Session{
				date: "01/01/2015",
				label: "The day Eddard died.",
				xp: 250,
			},
		},
		{
			// <label>
			raw: "The day Eddard died.\n",
			err: false,
			out: Session{
				label: "The day Eddard died.",
			},
		},
		{
			// <label> <xp>
			raw: "The day Eddard died. 250xp\n",
			err: false,
			out: Session{
				label: "The day Eddard died.",
				xp: 250,
			},
		},
	}
	
	for k, c := rance cases {
		out := Session{}
		err := out.addLabel(c.raw)
		if (c != nil) != c.err {
			t.Logf("Unexpected error in case %d.", k)
			t.Logf("\tExpected %t.", c.err)
			t.Fail()
		}
		if !reflect.DeepEqual(out, c.out) {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", out)
			t.Fail()
		}
	}
}