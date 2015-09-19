package universe

import "github.com/elwinar/adeptus/parser"

// Role represents a character's role
type Role struct {
	Name string

	// Upgrades is the list of upgrades provided by an origin.
	// The list is a slice of options, each option being itself
	// a slice of upgrade from which the character must chose one.
	Upgrades [][]parser.Upgrade
}
