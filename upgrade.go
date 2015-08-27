package adeptus

import (
	"fmt"
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
		
		// matches (250xp)
		if strings.HasPrefix(field, "(") && strings.HasSuffix(field, "xp)") {
			cost = strings.TrimSuffix(strings.TrimPrefix(field, "("), "xp)")
			// TODO: should be integer
			if len(cost) == 0 {
				return upgrade, fmt.Errorf("Error on line %d: xp has no value. Expected number.", line)
			}
		}
		
		// matches 250xp
		if strings.HasPrefix(field, "xp") {
			cost = strings.TrimSuffix(field, "xp")
			// TODO: should be integer
			if len(cost) == 0 {
				return upgrade, fmt.Errorf("Error on line %d: xp has no value. Expected number.", line)
			}
		}
		
		if len(cost) != 0 {
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
