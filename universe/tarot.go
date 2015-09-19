package universe

import "github.com/elwinar/adeptus/parser"

// Tarot is a character's trait providing him alterations.
type Tarot struct {
	Name string

	// Min is the lowest value of the tarot range.
	Min int

	// Max is the highest value of the tarot range.
	Max int

	// Upgrades is the list of upgrades provided by a tarot.
	// The list is a slice of options, each option being itself
	// a slice of upgrade from which the character must chose one.
	Upgrades [][]parser.Upgrade
}
