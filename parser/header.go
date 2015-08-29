package parser

import (
	"fmt"
	"strings"
)

// Header is the first block of the sheet, and define the character with its
// name, origin, etc.
type Header struct {
	Name       string
	Origin     string
	Background string
	Role       string
	Tarot      string
}

// ParseHeader generate a Header from a block of lines
func parseHeader(block []line) (Header, error) {
	// Initialize the values to find
	var name, origin, background, role, tarot string

	for _, line := range block {
		// Parse the field as a key and value
		fields := strings.Split(line.Text, ":")
		if len(fields) != 2 {
			return Header{}, fmt.Errorf("line %d: expected pair key:value", line.Number)
		}
		key := strings.TrimSpace(strings.ToLower(fields[0]))
		value := strings.TrimSpace(fields[1])

		// Check key:value
		switch key {
		case "name":
			name = value
		case "origin":
			origin = value
		case "background":
			background = value
		case "role":
			role = value
		case "tarot":
			tarot = value
		default:
			return Header{}, fmt.Errorf("line %d: unknown key \"%s\"", line.Number, key)
		}
	}

	// Check that every field is defined
	if name == "" {
		return Header{}, fmt.Errorf("undefined character's name")
	}

	if origin == "" {
		return Header{}, fmt.Errorf("undefined character's origin")
	}

	if background == "" {
		return Header{}, fmt.Errorf("undefined character's background")
	}

	if role == "" {
		return Header{}, fmt.Errorf("undefined character's role")
	}

	if tarot == "" {
		return Header{}, fmt.Errorf("undefined character's tarot")
	}

	return Header{
		Name:       name,
		Origin:     origin,
		Background: background,
		Role:       role,
		Tarot:      tarot,
	}, nil
}
