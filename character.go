package main

import (
	"adeptus/parser"
	"adeptus/universe"
	"fmt"
	"os"
)

// Character is the type representing a role playing character
type Character struct {
	Name       string
	Universe   universe.Universe
	Origin     universe.Origin
	Background universe.Background
	Role       universe.Role
	Tarot      universe.Tarot
}

// NewCharacter creates a new character given a sheet
func NewCharacter(s parser.Sheet) (*Character, error) {

	// Retrieve name
	if len(s.Header.Name) == 0 {
		return nil, NewError(UndefinedName)
	}
	name := s.Header.Name

	// Open and parse universe given the sheet's universe
	if s.Header.Universe.Label == nil {
		return nil, NewError(UndefinedUniverse)
	}

	reader, err := os.Open(fmt.Sprintf("samples/%s.json", *s.Header.Universe.Label))
	if err != nil {
		return nil, NewError(NotFoundUniverse, err)
	}

	data, err := universe.ParseUniverse(reader)
	if err != nil {
		return nil, NewError(InvalidUniverse, err)
	}

	// Retrieve origin
	if s.Header.Origin.Label == nil {
		return nil, NewError(UndefinedOrigin)
	}
	origin := universe.NewOrigin(*s.Header.Origin.Label)

	// Retrieve background
	if s.Header.Background.Label == nil {
		return nil, NewError(UndefinedBackground)
	}
	background := universe.NewBackground(*s.Header.Background.Label)

	// Retrieve role
	if s.Header.Role.Label == nil {
		return nil, NewError(UndefinedRole)
	}
	role := universe.NewRole(*s.Header.Role.Label)

	// Retrieve tarot
	if s.Header.Tarot.Label == nil {
		return nil, NewError(UndefinedTarot)
	}
	tarot := universe.NewTarot(*s.Header.Tarot.Label)

	return &Character{
		Name:       name,
		Universe:   data,
		Origin:     origin,
		Background: background,
		Role:       role,
		Tarot:      tarot,
	}, nil
}

// Debug prints the current values of the character
func (c Character) Debug() {
	fmt.Printf("Name		%s\n", c.Name)
	fmt.Printf("Origin		%s\n", c.Origin.Name)
	fmt.Printf("Background	%s\n", c.Background.Name)
	fmt.Printf("Role		%s\n", c.Role.Name)
	fmt.Printf("Tarot		%s\n", c.Tarot.Name)
}
