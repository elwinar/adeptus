package parser

import "strings"

// Header is the first block of the sheet, and define the character with its
// name, origin, etc.
type Header struct {
	Name       string
	Metas	   map[string][]Meta
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
	metas := make(map[string][]Meta)

	for _, line := range block {
		// Parse the field as a key and value.
		fields := strings.Split(line.Text, ":")
		if len(fields) != 2 {
			return Header{}, NewError(line.Number, InvalidKeyValuePair)
		}
		key := strings.ToLower(strings.TrimSpace(strings.ToLower(fields[0])))
		value := strings.TrimSpace(fields[1])
		
		// Check key is not empty
		if len(key) == 0 {
			return Header{}, NewError(line.Number, EmptyMetaKey)
		}
		
		// Check value is not empty
		if len(value) == 0 {
			return Header{}, NewError(line.Number, EmptyMetaValue)
		}

		// Retrieve the name.
		if key == "name" {
			name = value
			continue
		}
		
		// Check the meta is unique.
		_, found := metas[key]
		if found {
			return Header{}, NewError(line.Number, DuplicateMeta)
		}
		
		// Retrieve coma separated values.
		metas[key] = []Meta{}
		splits := strings.Split(value, ",")
		for _, s := range splits {
				meta, err := NewMeta(strings.TrimSpace(s))
				if err != nil {
					return Header{}, NewError(line.Number, InvalidOptions)
				}
				metas[key] = append(metas[key], meta)
		}
	}

	return Header{
		Name:       name,
		Metas:		metas,
	}, nil
}
