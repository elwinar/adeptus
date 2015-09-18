package parser

import "strings"

// Meta is a header and a collection of associated options
type Meta struct {
	Label   *string
	Options []string
}

// NewMeta constructs a new meta given a label
func NewMeta(label string) Meta {
	return Meta{
		Label: &label,
	}
}

// Header is the first block of the sheet, and define the character with its
// name, origin, etc.
type Header struct {
<<<<<<< HEAD
	Name       string
	Origin     Meta
	Background Meta
	Role       Meta
	Tarot      Meta
	Universe   Meta
}

// ParseHeader generate a Header from a block of lines. The block must not be
// empty.
func parseHeader(block []line) (Header, error) {
	// Check the block is non-empty
	if len(block) == 0 {
		panic("empty block")
	}

	// Initialize the values to find
	var name string
	var origin, background, role, tarot, universe Meta

	for _, line := range block {
		// Parse the field as a key and value
		fields := strings.Split(line.Text, ":")
		if len(fields) != 2 {
			return Header{}, NewError(line.Number, InvalidKeyValuePair)
		}
		key := strings.TrimSpace(strings.ToLower(fields[0]))
		value := strings.TrimSpace(fields[1])

		// Check key:value
		switch key {
		case "name":
			name = value
		case "origin":
			origin = NewMeta(value)
		case "background":
			background = NewMeta(value)
		case "role":
			role = NewMeta(value)
		case "tarot":
			tarot = NewMeta(value)
		case "universe":
			universe = NewMeta(value)
		default:
			return Header{}, NewError(line.Number, UnknownKey)
		}
	}

	return Header{
		Name:       name,
		Origin:     origin,
		Background: background,
		Role:       role,
		Tarot:      tarot,
		Universe:   universe,
	}, nil
}
