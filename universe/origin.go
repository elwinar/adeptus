package universe

import "github.com/elwinar/adeptus/parser"

// Origin represent a character's origin.
// It provide a list of upgrades to apply upon
// choice of this origin by the character.
type Origin struct {
	Name string

	// Upgrades is the list of upgrades provided by an origin.
	// The list is a slice of options, each option being itself
	// a slice of upgrade from which the character must chose one.
	Upgrades [][]parser.Upgrade
}

// NewOrigin returns the origin associated to the name
func NewOrigin(name string) Origin {
	return Origin{
		Name: name,
	}
}
