package adeptus

import (
	"fmt"
	"strings"
)

type Header struct {
	Name       string
	Origin     string
	Background string
	Role       string
	Tarot      string
}

// ParseHeader generate a Header from a block of lines
func ParseHeader(block []Line) (Header, error) {

	header := Header{}
	for _, line := range block {
		fields := strings.SplitN(line.Text, ":", 2)
		if len(fields) < 2 {
			return Header{}, fmt.Errorf("Error on line %d: Expected pair key:value.", line.Number)
		}

		// Check key:value
		key := strings.TrimSpace(strings.ToLower(fields[0]))
		value := strings.TrimSpace(fields[1])
		switch key {
		case "name":
			header.Name = value
		case "origin":
			header.Origin = value
		case "background":
			header.Background = value
		case "role":
			header.Role = value
		case "tarot":
			header.Tarot = value
		default:
			return Header{}, fmt.Errorf("Error on line %d: Undefined key: \"%s\".", line.Number, key)
		}
	}

	// every field must be informed
	if header.Name == "" {
		return Header{}, fmt.Errorf("Undefined character's name. Set name:<name> In header.")
	}
	if header.Origin == "" {
		return Header{}, fmt.Errorf("Undefined character's origin. Set origin:<origin> In header.")
	}
	if header.Background == "" {
		return Header{}, fmt.Errorf("Undefined character's background. Set background:<background> In header.")
	}
	if header.Role == "" {
		return Header{}, fmt.Errorf("Undefined character's role. Set role:<role> In header.")
	}
	if header.Tarot == "" {
		return Header{}, fmt.Errorf("Undefined character's tarot. Set tarot:<tarot> In header.")
	}

	return header, nil
}
