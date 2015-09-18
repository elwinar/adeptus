package universe

import "adeptus/parser"

type Tarot struct {
	Name string
	Upgrades [][]parser.Upgrade
}

// NewTarot returns the background associated to the name
func NewTarot(name string) Tarot {
		return Tarot{
			Name: name,
		}
}
