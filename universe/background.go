package universe

import "adeptus/parser"

type Background struct {
	Name string
	Upgrades [][]parser.Upgrade
}

// NewBackground returns the background associated to the name
func NewBackground(name string) Background {
		return Background{
			Name: name,
		}
}
