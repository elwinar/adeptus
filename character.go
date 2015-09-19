package main

import (
	"fmt"
	"strconv"

	"github.com/elwinar/adeptus/parser"
	"github.com/elwinar/adeptus/universe"
)

// Character is the type representing a role playing character
type Character struct {
	Name       string
	Origin     universe.Origin
	Background universe.Background
	Role       universe.Role
	Tarot      universe.Tarot
}

// NewCharacter creates a new character given a sheet
func NewCharacter(u universe.Universe, s parser.Sheet) (*Character, error) {

	var h parser.Header = s.Header

	var name string
	var origin universe.Origin
	var background universe.Background
	var role universe.Role
	var tarot universe.Tarot
	var found bool

	// Retrieve the name from the sheet header.
	if len(h.Name) == 0 {
		return nil, fmt.Errorf("empty name")
	}
	name = h.Name

	// Retrieve the origin from the universe.
	if h.Origin == nil {
		return nil, fmt.Errorf("unspecified origin")
	}
	origin, found = u.FindOrigin(h.Origin.Label)
	if !found {
		return nil, fmt.Errorf("origin %s not found", h.Origin.Label)
	}

	// Retrieve the background from the universe.
	if h.Background == nil {
		return nil, fmt.Errorf("unspecified background")
	}
	background, found = u.FindBackground(h.Background.Label)
	if !found {
		return nil, fmt.Errorf("background %s not found", h.Background.Label)
	}

	// Retrieve the role from the universe.
	if h.Role == nil {
		return nil, fmt.Errorf("unspecified role")
	}
	role, found = u.FindRole(h.Role.Label)
	if !found {
		return nil, fmt.Errorf("role %s not found", h.Role.Label)
	}

	// Retrieve the tarot from the universe.
	if h.Tarot == nil {
		return nil, fmt.Errorf("unspecified tarot")
	}

	dice, err := strconv.Atoi(h.Tarot.Label)
	if err == nil {
		tarot, found = u.FindTarotByDice(dice)
	} else {
		tarot, found = u.FindTarot(h.Tarot.Label)
	}
	if !found {
		return nil, fmt.Errorf("tarot %s not found", h.Tarot.Label)
	}

	return &Character{
		Name:       name,
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
