package parser

import "strings"

// Header is the first block of the sheet, and define the character with its
// name, origin, etc.
type Header struct {
	Name       string
	Origin     *Meta
	Background *Meta
	Role       *Meta
	Tarot      *Meta
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
	var origin, background, role, tarot *Meta

	for _, line := range block {
		// Parse the field as a key and value.
		fields := strings.Split(line.Text, ":")
		if len(fields) != 2 {
			return Header{}, NewError(line.Number, InvalidKeyValuePair)
		}
		key := strings.TrimSpace(strings.ToLower(fields[0]))
		value := strings.TrimSpace(fields[1])

		// Check the key is knowned.
		if !in(key, []string{"name", "origin", "background", "role", "tarot"}) {
			return Header{}, NewError(line.Number, UnknownKey)
		}

		// Retrieve the name.
		if key == "name" {
			name = value
			continue
		}

		// Create new meta.
		m, err := NewMeta(value)
		if err != nil {
			return Header{}, NewError(line.Number, InvalidOptions)
		}

		// Associate the proper key to the meta.
		switch key {
		case "origin":
			if origin != nil {
				return Header{}, NewError(line.Number, DuplicateMeta)
			}
			origin = &m
		case "background":
			if background != nil {
				return Header{}, NewError(line.Number, DuplicateMeta)
			}
			background = &m
		case "role":
			if role != nil {
				return Header{}, NewError(line.Number, DuplicateMeta)
			}
			role = &m
		case "tarot":
			if tarot != nil {
				return Header{}, NewError(line.Number, DuplicateMeta)
			}
			tarot = &m
		}
	}

	return Header{
		Name:       name,
		Origin:     origin,
		Background: background,
		Role:       role,
		Tarot:      tarot,
	}, nil
}
