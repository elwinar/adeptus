package main

import (
		"adeptus/universe"
		"adeptus/parser"
)

// Character is the type representing a role playing character
type Character struct{
		Name 		string
		Universe	universe.Universe
		Origin		universe.Origin
		Background	universe.Background
		Role		universe.Role
		Tarot		universe.Tarot
}

// NewCharacter creates a new character given a sheet
func NewCharacter(sheet parser.Sheet) *Character {
	c := &Character{}
	return c
}