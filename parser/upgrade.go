package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Upgrade describe one upgrade applied to the character. The mark indicate how
// to handle the upgrade considéring future upgrades. The cost is optionnal in
// certain cases.
type Upgrade struct {
	Mark string
	Name string
	Cost *int
}

// parseUpgrade generate an upgrade from a raw line
func parseUpgrade(line line) (Upgrade, error) {
	// Get the fields of the line
	fields := strings.Fields(line.Text)

	// The minimum number of fields is 2
	if len(fields) < 2 {
		return Upgrade{}, fmt.Errorf("line %d: expected at least a mark and a label", line.Number)
	}

	// Check that the mark is a valid one
	if !in(fields[0], []string{"*", "+", "-"}) {
		return Upgrade{}, fmt.Errorf("line %d: \"%s\" is not a valid mark (\"*\", \"+\", \"-\")", line.Number, fields[0])
	}

	// Set the upgrade mark
	mark := fields[0]
	fields = fields[1:]

	// Check if a field seems to be a cost field
	var cost *int
	for i, field := range fields {
		// If one end has the brackets but not the other, that's an error:
		// brackets does by pairs, and are forbidden in the title
		if strings.HasPrefix(field, "[") != strings.HasSuffix(field, "]") {
			return Upgrade{}, fmt.Errorf("line %d: brackets [] must open-close and contain no blank", line.Number)
		}

		// If the brackets are absents, that's not a cost, so skip the field.
		// Note that as both ends have brackets (or not), we just need to test
		// one of them.
		if !strings.HasPrefix(field, "[") {
			break
		}

		// Check position of the cost
		if i != 0 && i != len(fields)-1 {
			return Upgrade{}, fmt.Errorf("line %d: cost must be in second or last position of the line", line.Number)
		}

		// Trim the field to get the raw cost
		raw := strings.Trim(field, "[]")

		// Parse the cost
		c, err := strconv.Atoi(raw)
		if err != nil {
			return Upgrade{}, fmt.Errorf("line %d: expected integer, got \"%s\"", line.Number, raw)
		}
		cost = &c

		// Remove the field from the slice
		fields = append(fields[:i], fields[i+1:]...)
	}

	// The remaining line is the name of the upgrade
	if len(fields) == 0 {
		return Upgrade{}, fmt.Errorf("line %d: name should not be empty", line.Number)
	}

	return Upgrade{
		Mark: mark,
		Name: strings.Join(fields, " "),
		Cost: cost,
	}, nil
}
