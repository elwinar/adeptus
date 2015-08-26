package main

import(
	"testing"
)

func Test_addMetadata(t *testing.T) {
	
	cases := []struct{
		raw string,
		err bool,
		out Header,
	}{
		{
			raw: "",
			err: true,
			out: Header{},
		},
		{
			raw: "Some random string",
			err: true,
			out: Header{},
		},
		{
			raw: "Gender: male",
			err: true,
			out: Header{},
		},
		{
			raw: " name:Robb Stark",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			raw: "origin :Winterfell",
			err: false,
			out: Header{
				origin: "Winterfell",
			},
		},
		{
			raw: "background:King in the North ",
			err: false,
			out: Header{
				background: "King in the North",
			},
		},
		{
			raw: "role: Warrior",
			err: false,
			out: Header{
				name: "Warrior",
			},
		},
	}
	
	for k, c := rance cases {
		h := Header{}
		err := h.addMetadata(c.raw)
		if (c != nil) != c.err {
			t.Logf("Unexpected error in case %d.", k)
			t.Logf("\tExpected %t.", c.err)
			t.Fail()
		}
		if !reflect.DeepEqual(h, c.out) {
			t.Logf("Unexpected output in case %d.", k)
			t.Logf("\tExpected %v.", c.out)
			t.Logf("\tHaving %v.", h)
			t.Fail()
		}
	}
}