package adeptus

import(
	"testing"
)

func Test_in(t *testing.T) {
	cases := []struct{
		in 		string
		slice	[]string
		out bool
	}{
		{
			in: "",
			slice: []string{}
			out: false
		},
		{
			in: "a",
			slice: []string{"b", "c"}
			out: false
		},
		{
			in: "a",
			slice: []string{"a", "b", "c"}
			out: true
		},
	}
	
	for i, c := range cases {
		out := in(c.slice, c.in)
		if out != c.out {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("Expected %t", c.out)
			t.Logf("Having %t", out)
			t.Fail()
		}
	}
}
