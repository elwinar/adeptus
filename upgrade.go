package adeptus

import (
	"fmt"
	"strconv"
	"strings"
)

type Upgrade interface {
}

type RawUpgrade struct {
	mark       string
	name       string
	cost       int
	customCost bool
}

// parseUpgrade generate an upgrade from a raw line
func parseUpgrade(line Line) (RawUpgrade, error) {
	upgrade := RawUpgrade{}

	// Get the fields of the line
	fields := strings.Fields(line.Text)

	// The minimum number of fields is 2
	if len(fields) < 2 {
		return upgrade, fmt.Errorf("Error on line %d: expected at least mark and label.", line.Number)
	}

	// Check that the mark is a valid one
	if !in(fields[0], []string{"*", "+", "-"}) {
		return upgrade, fmt.Errorf("Error on line %d: \"%s\" is not a valid mark (\"*\", \"+\", \"-\").", line.Number, fields[0])
	}

	// Set the upgrade mark
	mark := fields[0]
	fields = fields[1:]

	// Check if a field seems to be a cost field
	var customCost bool = false
	var cost int
	var err error
	for i, field := range fields {

		if strings.HasPrefix(field, "[") || strings.HasSuffix(field, "]") {

			// Check that the field has both brackets. If only one bracket is present, there is an error
			if strings.HasPrefix(field, "[") != strings.HasSuffix(field, "]") {
				return upgrade, fmt.Errorf("Error on line %d: brackets [] must open-close and contain no blank.", line.Number)
			}

			// Check position of xp
			if i != 0 && i != len(fields)-1 {
				return upgrade, fmt.Errorf("Error on line %d: experience must be in second or last position of the line.", line.Number)
			}

			// Check value of xp
			xp := strings.TrimSuffix(strings.TrimPrefix(field, "["), "]")
			cost, err = strconv.Atoi(xp)
			if err != nil || len(xp) == 0 {
				return upgrade, fmt.Errorf("Error on line %d: expected number, \"%s\" is no numeric value.", line.Number, xp)
			}

			// remove xp from field slice
			fields = append(fields[:i], fields[i+1:]...)
			customCost = true
			break
		}
	}

	// The remaining line is the name of the upgrade
	if len(fields) == 0 {
		return upgrade, fmt.Errorf("Error on line %d: name should not be empty.", line.Number)
	}

	// Set the upgrade attributes at the end to return empty upgrade in case of error
	upgrade.customCost = customCost
	upgrade.mark = mark
	upgrade.cost = cost
	upgrade.name = strings.Join(fields, " ")

	return upgrade, nil
}
