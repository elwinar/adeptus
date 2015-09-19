package universe

import "github.com/elwinar/adeptus/parser"

// Background represent a character's background
type Background struct {
	Name string

	// Upgrades is the list of upgrades provided by a background.
	// The list is a slice of options, each option being itself
	// a slice of upgrade from which the character must chose one.
	Upgrades [][]parser.Upgrade
}
