package universe

import "github.com/elwinar/adeptus/parser"

// History represents a part of a character's history responsible for changes
// on it's definition. Typically, it represents the origin, background, role and 
// tarot of the character
type History interface{
	GetName() string
	GetUpgrades() [][]parser.Upgrade
}

// historyTrait is the struct inherited by any struct implementing the History interface
type history struct{
	Name string
	Upgrades [][]parser.Upgrade
}

// GetName returns the name of the history
func (h history) GetName() string {
	return h.Name
}


// GetUpgrades returns the list of upgrades provided by a background.
// The list is a slice of options, each option being itself
// a slice of upgrade from which the character must chose one.
func (h history) GetUpgrades() [][]parser.Upgrade {
	return h.Upgrades
}