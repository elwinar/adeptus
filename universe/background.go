package universe

import "adeptus/parser"

// Background represent a character's background
type Background struct {
	Name     string
	Upgrades [][]parser.Upgrade
}

// NewBackground returns the background associated to the name
func NewBackground(name string) Background {
	return Background{
		Name: name,
	}
}
