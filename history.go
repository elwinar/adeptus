package main

// History represents an element providing traits to a character.
// Background and role are History.
type History struct {
	Type     string
	Name     string
	Upgrades [][]Upgrade
}
