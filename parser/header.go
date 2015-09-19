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
		// Parse the field as a key and value
		fields := strings.Split(line.Text, ":")
		if len(fields) != 2 {
			return Header{}, NewError(line.Number, InvalidKeyValuePair)
		}
		key := strings.TrimSpace(strings.ToLower(fields[0]))
		value := strings.TrimSpace(fields[1])

		var err error

		// Parse the value depending on the key
		switch key {
		case "name":
			name = value
		case "origin":
			origin, err = NewMeta(value)
		case "background":
			background, err = NewMeta(value)
		case "role":
			role, err = NewMeta(value)
		case "tarot":
			tarot, err = NewMeta(value)
		// If the key doesn't exists, return an error
		default:
			return Header{}, NewError(line.Number, UnknownKey)
		}

		// If there was an error parsing the value, return it
		if err != nil {
			return Header{}, NewError(line.Number, InvalidMeta)
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
