package universe

import "adeptus/parser"

type Origin struct {
	Name string
	Upgrades [][]parser.Upgrade
}

// NewOrigin returns the origin associated to the name
func NewOrigin(name string) Origin {
		return Origin{}
}