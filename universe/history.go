package universe

import "github.com/elwinar/adeptus/parser"

// History represents an element providing traits to a character.
// Background and role are History.
type History struct {
	Name     string
	Upgrades [][]parser.Upgrade
}
