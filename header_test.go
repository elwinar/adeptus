package main

import(
	"testing"
)

func Test_Header_addMetadata(t *testing.T) {
	
	cases := []struct{
		raw string,
		err bool,
		out Header,
	}{
		// errors
		{
			// nil input
			raw: "",
			err: true,
			out: Header{},
		},
		{
			// no value
			raw: "Some random string\n",
			err: true,
			out: Header{},
		},
		{
			// unrecognized key
			raw: "Gender: male\n",
			err: true,
			out: Header{},
		},
		// parsing
		{
			// <key>:<value>
			raw: "name:Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <Key>:<value>
			raw: "Name:Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <KEY>:<value>
			raw: "NAME:Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <key>:( )<value>
			raw: "name: Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <key>:( )*<value>
			raw: "name:     Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <key>:(\t)<value>
			raw: "name:\tRobb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <key>:(\t)*<value>
			raw: "name:\t\t\tRobb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <key>( ):<value>
			raw: "name :Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <key>( )*:<value>
			raw: "name   :Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <key>(\t):<value>
			raw: "name\t:Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// <key>(\t)*:<value>
			raw: "name\t\t\t:Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		// keys
		{
			// name
			raw: "name: Robb Stark\n",
			err: false,
			out: Header{
				name: "Robb Stark",
			},
		},
		{
			// origin
			raw: "origin: Winterfell\n",
			err: false,
			out: Header{
				origin: "Winterfell",
			},
		},
		{
			// background
			raw: "background: King in the North\n",
			err: false,
			out: Header{
				background: "King in the North",
			},
		},
		{
			// role
			raw: "role: Warrior\n",
			err: false,
			out: Header{
				role: "Warrior",
			},
		},
	}
	
	for k, c := rance cases {
		out := Header{}
		err := out.addMetadata(c.raw)
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