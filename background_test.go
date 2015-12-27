package main

import (
	"reflect"
	"testing"
)

func Test_Background_Apply(t *testing.T) {

	universe := Universe{
		Backgrounds: map[string][]Background{
			"origin": []Background{
				{
					Type: "origin",
					Name: "france",
					Upgrades: []string{
						"awesomeness",
						"blacchusness",
						"WS +5",
					},
				},
			},
		},
		Aptitudes: []Aptitude{
			Aptitude("awesomeness"),
		},
		Talents: []Talent{
			{Name: "blacchusness"},
		},
		Characteristics: []Characteristic{
			{Name: "WS"},
		},
	}

	for i, c := range []struct {
		background Background
		character  Character
		out        Character
		err        bool
		code       ErrorCode
	}{
		{
			background: Background{
				Type: "origin",
				Name: "france",
				Upgrades: []string{
					"awesomeness",
					"blacchusness",
					"WS +5",
				},
			},
			character: Character{
				Backgrounds:     map[string]Background{},
				Characteristics: map[string]Characteristic{},
				Talents:         map[string]Talent{},
				Aptitudes:       map[string]Aptitude{},
			},
			out: Character{
				Backgrounds: map[string]Background{
					"france": Background{
						Type: "origin",
						Name: "france",
						Upgrades: []string{
							"awesomeness",
							"blacchusness",
							"WS +5",
						},
					},
				},
				Talents: map[string]Talent{
					"blacchusness": Talent{Name: "blacchusness", Value: 1},
				},
				Characteristics: map[string]Characteristic{
					"WS": Characteristic{
						Name:  "WS",
						Value: 5,
						Tier:  0,
					},
				},
				Aptitudes: map[string]Aptitude{
					"awesomeness": Aptitude("awesomeness"),
				},
				History: []Upgrade{
					Upgrade{Mark: MarkApply, Name: "awesomeness", Cost: IntP(0)},
					Upgrade{Mark: MarkApply, Name: "blacchusness", Cost: IntP(0)},
					Upgrade{Mark: MarkSpecial, Name: "WS +5", Cost: IntP(0)},
				},
			},
			err: false,
		},
		{
			background: Background{
				Type: "origin",
				Name: "france",
				Upgrades: []string{
					"awesomeness",
					"blacchusness",
					"WS +5",
				},
			},
			character: Character{
				Backgrounds:     map[string]Background{},
				Characteristics: map[string]Characteristic{},
				Talents: map[string]Talent{
					"blacchusness": Talent{Name: "blacchusness", Value: 1},
				},
				Aptitudes: map[string]Aptitude{},
			},
			out:  Character{},
			err:  true,
			code: DuplicateUpgrade,
		},
	} {
		err := c.background.Apply(&c.character, universe)
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
				t.Logf("	Expected %s", NewError(c.code))
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
