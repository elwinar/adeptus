package adeptus

import (
	"fmt"
	"strconv"
	"strings"
)

type Upgrade interface {
}

type RawUpgrade struct {
	mark string
	name string
	cost string
	line int
}

// ParseUpgrade generate an upgrade from a raw line
func ParseUpgrade(line int, raw string) (RawUpgrade, error) {
	upgrade := RawUpgrade{
		line: line,
	}

	// Get the fields of the line
	fields := strings.Fields(raw)

	// The minimum number of fields is 2
	if len(fields) < 2 {
		return upgrade, fmt.Errorf("Error on line %d: expected at least mark and label.", line)
	}

	// Check that the mark is a valid one
	if !in(fields[0], []string{"*", "+", "-"}) {
		return upgrade, fmt.Errorf("Error on line %d: \"%s\" is not a valid mark (\"*\", \"+\", \"-\").", line, fields[0])
	}

	// Set the upgrade mark
	mark := fields[0]
	fields = fields[1:]

	// Check if a field seems to be a cost field
	var cost string
	for i, field := range fields {
		
		if strings.HasPrefix(field, "[") || strings.HasSuffix(field, "]") {
		
			// Check that the field has both brackets. If only one bracket is present, there is an error
			if strings.HasPrefix(field, "[") != strings.HasSuffix(field, "]") {
				return upgrade, fmt.Errorf("Error on line %d: brackets [] must open-close and contain no blank.", line)
			}
		
			// Check position of xp
			if i == 0 || i == len(fields) - 1 {
				return upgrade, fmt.Errorf("Error on line %d: experience must be after mark or at the end of line.", line)
			}
			
			// Check value of xp
			cost = strings.TrimSuffix(strings.TrimPrefix(field, "["), "]")
			_, err := strconv.Atoi(cost)
			if err || len(cost) == 0 {
				return upgrade, fmt.Errorf("Error on line %d: expected number, \"%s\" is no numeric value.", line, cost)
			}
			
			// remove xp from field slice
			fields = append(fields[:i], fields[i+1:]...)
			break
		}
	}

	// The remaining line is the name of the upgrade
	if len(fields) == 0 {
		return upgrade, fmt.Errorf("Error on line %d: name should not be empty.", line)
	}

	// Set the upgrade attributes at the end to return empty upgrade in case of error
	upgrade.mark = mark
	upgrade.cost = cost
	upgrade.name = strings.Join(fields, " ")

	return upgrade, nil
}
