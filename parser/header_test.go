package parser

import (
	"reflect"
	"testing"
)

func Test_ParseHeader(t *testing.T) {
	cases := []struct {
		in  []string
		out Header
		err bool
	}{
		{
			in:  []string{},
			out: Header{},
			err: true,
		},
		{
			in: []string{
				"fail",
			},
			out: Header{},
			err: true,
		},
		{
			in: []string{
				":",
			},
			out: Header{},
			err: true,
		},
		{
			in: []string{
				"fail:fail",
			},
			out: Header{},
			err: true,
		},
		{
			in: []string{
				"name: fail",
			},
			out: Header{},
			err: true,
		},
		{
			in: []string{
				"name: fail",
				"origin: fail",
			},
			out: Header{},
			err: true,
		},
		{
			in: []string{
				"name: fail",
				"origin: fail",
				"background: fail",
			},
			out: Header{},
			err: true,
		},
		{
			in: []string{
				"name: fail",
				"origin: fail",
				"background: fail",
				"role: fail",
			},
			out: Header{},
			err: true,
		},
		{
			in: []string{
				"rolE: successful role",
				"name: successful name",
				"tarot: successful tarot",
				"background: successful background",
				"origin: successful origin",
			},
			out: Header{
				Name:       "successful name",
				Origin:     "successful origin",
				Background: "successful background",
				Role:       "successful role",
				Tarot:      "successful tarot",
			},
			err: false,
		},
		{
			in: []string{
				"role	: successful role",
				"name	: successful name",
				"tarot	: successful tarot",
				"background	: successful background",
				"origin	: successful origin",
			},
			out: Header{
				Name:       "successful name",
				Origin:     "successful origin",
				Background: "successful background",
				Role:       "successful role",
				Tarot:      "successful tarot",
			},
			err: false,
		},
		{
			in: []string{
				"role: 	successful role",
				"name: 	successful name",
				"tarot: 	successful tarot",
				"background: 	successful background",
				"origin: 	successful origin",
			},
			out: Header{
				Name:       "successful name",
				Origin:     "successful origin",
				Background: "successful background",
				Role:       "successful role",
				Tarot:      "successful tarot",
			},
			err: false,
		},
		{
			in: []string{
				"	role: successful role",
				"	name: successful name",
				"	tarot: successful tarot",
				"	background: successful background",
				"	origin: successful origin",
			},
			out: Header{
				Name:       "successful name",
				Origin:     "successful origin",
				Background: "successful background",
				Role:       "successful role",
				Tarot:      "successful tarot",
			},
			err: false,
		},
	}

	for i, c := range cases {
		in := []line{}
		for number, text := range c.in {
			in = append(in, newLine(text, number))
		}

		out, err := parseHeader(in)
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
