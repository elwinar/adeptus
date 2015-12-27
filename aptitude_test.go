package main

import (
	"reflect"
	"testing"
)

func Test_Aptitude_Cost(t *testing.T) {
	in := Aptitude("awesomeness")

	out, err := in.Cost(Universe{}, Character{})
	expected := 0

	if err != nil {
		t.Logf("Expected error")
		t.Fail()
	}

	if out != expected {
		t.Logf("Unexpected output")
		t.Logf("	Expected %v", expected)
		t.Logf("	Having %v", out)
		t.Fail()
	}
}

func Test_Aptitude_Apply(t *testing.T) {

	a := Aptitude("awesomeness")

	for i, c := range []struct {
		upgrade   Upgrade
		character Character
		out       Character
		err       bool
		code      ErrorCode
	}{
		{
			upgrade: Upgrade{
				Mark: MarkSpecial,
				Line: 0,
				Cost: nil,
				Name: "awesomeness",
			},
			character: Character{
				Aptitudes: map[string]Aptitude{},
			},
			out:  Character{},
			err:  true,
			code: ForbidenUpgradeMark,
		},
		{
			upgrade: Upgrade{
				Mark: MarkRevert,
				Line: 0,
				Cost: nil,
				Name: "awesomeness",
			},
			character: Character{
				Aptitudes: map[string]Aptitude{},
			},
			out:  Character{},
			err:  true,
			code: ForbidenUpgradeLoss,
		},
		{
			upgrade: Upgrade{
				Mark: MarkRevert,
				Line: 0,
				Cost: nil,
				Name: "awesomeness",
			},
			character: Character{
				Aptitudes: map[string]Aptitude{
					"awesomeness": Aptitude("awesomeness"),
				},
			},
			out: Character{
				Aptitudes: map[string]Aptitude{},
			},
			err: false,
		},
		{
			upgrade: Upgrade{
				Mark: MarkApply,
				Line: 0,
				Cost: nil,
				Name: "awesomeness",
			},
			character: Character{
				Aptitudes: map[string]Aptitude{},
			},
			out: Character{
				Aptitudes: map[string]Aptitude{
					"awesomeness": Aptitude("awesomeness"),
				},
			},
			err: false,
		},
	} {
		err := a.Apply(&c.character, c.upgrade)
		if (err != nil) != c.err {
			if c.err {
				t.Logf("Expected error on case %d", i+1)
				t.Fail()
			} else {
				t.Logf("Unexpected error on case %d: %s", i+1, err)
				t.Fail()
			}
		} else if err != nil {
			code := err.(Error).Code
			if c.code != code {
				t.Logf("Unexpected error on case %d:", i+1)
				t.Logf("	Expected %s", NewError(c.code, c.upgrade.Line))
				t.Logf("	Having %s", err)
				t.Fail()
			}
		}

		if err != nil {
			continue
		}

		out := c.character
		if !reflect.DeepEqual(out, c.out) {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}
}
