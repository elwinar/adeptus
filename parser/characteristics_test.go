package parser

import (
	"reflect"
	"testing"
)

func Test_parseCharacteristics(t *testing.T) {
	cases := []struct {
		in    []string
		out   Characteristics
		err   bool
		panic bool
	}{
		{
			in:    []string{},
			out:   Characteristics{},
			err:   false,
			panic: true,
		},
		{
			in:    []string{
                            "",
                        },
			out:   Characteristics{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
                            " ",
                        },
			out:   Characteristics{},
			err:   true,
			panic: false,
		},
		{
			in: []string{
                            " 	 ",
                        },
			out:   Characteristics{},
			err:   true,
			panic: false,
		},
		{
			in:  []string{
                            "fail",
                        },
			out: Characteristics{},
			err: true,
			panic: false,
		},
		{
			in:  []string{
                            "x fail",
                        },
			out: Characteristics{},
			err: true,
			panic: false,
		},
		{
			in:  []string{
                            "* WP 25",
                        },
			out: Characteristics{},
			err: true,
			panic: false,
		},
		{
			in:  []string{
                            "WP +25",
                        },
			out: Characteristics{},
			err: true,
			panic: false,
		},
		{
			in:  []string{
                            "STR 25", 
                            "WP +25",
                        },
			out: Characteristics{},
			err: true,
			panic: false,
		},
		{
			in:  []string{
                            "WP 25",
                        },
			out: Characteristics{
                                Upgrade{
                                        Mark: "-",
                                        Name: "WP 25",
                                        Cost: nil,
                                },
                        },
			err: false,
			panic: false,
		},
		{
			in:  []string{
                            "STR 25", 
                            "WP 25",
                        },
			out: Characteristics{
                                Upgrade{
                                        Mark: "-",
                                        Name: "STR 25",
                                        Cost: nil,
                                },
                                Upgrade{
                                        Mark: "-",
                                        Name: "WP 25",
                                        Cost: nil,
                                },
                        },
			err: false,
			panic: false,
		},
	}

	for i, c := range cases {
		in := []line{}
		for number, text := range c.in {
			in = append(in, newLine(text, number))
		}

		out, err, panic := func() (out Characteristics, err error, panic bool) {
			defer func() {
				if e := recover(); e != nil {
					panic = true
				}
			}()
                        out = Characteristics{}

			out, err = parseCharacteristics(in)
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
