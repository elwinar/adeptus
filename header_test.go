package adeptus

import(
	"testing"
)

func Test_ParseHeader(t *testing.T) {
	cases := []struct{
		in []Line
		out	Header
		err	bool
	}{
		{
			in: []Line{},
			out: Header{},
			err: true
		},
		{
			in: []Line{
				Line{Text: "fail"},
			},
			out: Header{},
			err: true
		},
		{
			in: []Line{
				Line{Text: ":"},
			},
			out: Header{},
			err: true
		},
		{
			in: []Line{
				Line{Text: "fail:fail"},
			},
			out: Header{},
			err: true
		},
		{
			in: []Line{
				Line{Text: "NamE:success"},
			},
			out: Header{
				Name: "success"
			},
			err: false
		},
		{
			in: []Line{
				Line{Text: "	name	:	success	"},
			},
			out: Header{
				Name: "success"
			},
			err: false
		},
		{
			in: []Line{
				Line{Text: "role: successful role"},
				Line{Text: "name: successful name"},
				Line{Text: "tarot: successful tarot"},
				Line{Text: "background: successful background"},
				Line{Text: "origin: successful origin"},
			},
			out: Header{
				Name: "successful name",
				Origin: "successful origin",
				Background: "successful background",
				Role: "successful role",
				Tarot: "successful tarot",
			},
			err: false
		},
	}

	for i, c := range cases {
		out, err := ParseHeader(c.in)
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