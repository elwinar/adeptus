package universe

import "adeptus/parser"

// Tarot is a character's trait providing him alterations
type Tarot struct {
	Name     string
	Upgrades [][]parser.Upgrade
}

// NewTarot returns the background associated to the name
func NewTarot(name string) Tarot {
	return Tarot{
		Name: name,
	}
}
