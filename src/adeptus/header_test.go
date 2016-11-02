package main

import (
	"reflect"
	"testing"
)

func Test_ParseHeader(t *testing.T) {
	cases := []struct {
		in    []string
		out   Header
		err   bool
		panic bool
	}{
		{
			in:    []string{},
			out:   Header{},
			err:   false,
			panic: true,
		},
		{
			in: []string{
				"fail",
			},
			out:   Header{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				":",
			},
			out:   Header{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"a:",
			},
			out:   Header{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"name: success",
			},
			out: Header{
				Name:  "success",
				Metas: map[string][]Meta{},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"origin: fail(",
			},
			out:   Header{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
				"rolE: successful role",
				"name: successful name",
			},
			out: Header{
				Name: "successful name",
				Metas: map[string][]Meta{
					"role": {
						Meta{
							Label:   "successful role",
							Line:    1,
							Options: nil,
						},
					},
				},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"role	: successful role",
				"name	: successful name",
				"tarot	: successful tarot",
			},
			out: Header{
				Name: "successful name",
				Metas: map[string][]Meta{
					"role": {
						Meta{
							Label:   "successful role",
							Line:    1,
							Options: nil,
						},
					},
					"tarot": {
						Meta{
							Label:   "successful tarot",
							Line:    3,
							Options: nil,
						},
					},
				},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"role: 	successful role, second role",
				"name: 	successful name",
			},
			out: Header{
				Name: "successful name",
				Metas: map[string][]Meta{
					"role": {
						Meta{
							Label: "successful role",
							Line:  1,
						},
						Meta{
							Label: "second role",
							Line:  1,
						},
					},
				},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"	role: successful role",
				"	name: successful name",
			},
			out: Header{
				Name: "successful name",
				Metas: map[string][]Meta{
					"role": {
						Meta{
							Label: "successful role",
							Line:  1,
						},
					},
				},
			},
			err:   false,
			panic: false,
		},
		{
			in: []string{
				"	role: fail",
				"	role: fail",
			},
			out:   Header{},
			err:   true,
			panic: false,
		},
	}

	for i, c := range cases {
		in := []line{}
		for number, text := range c.in {
			in = append(in, newLine(text, number+1))
		}

		out, err, panic := func() (out Header, err error, panic bool) {
			defer func() {
				if e := recover(); e != nil {
					panic = true
				}
			}()

			out, err = parseHeader(in)
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
