package main

import (
	"strconv"
	"strings"
)

const (
	// MarkApply denotes that the upgrade must be applied. It will provide the upgrade to the character,
	// increase the value of the trait and increment its tier if able.
	MarkApply = "+"

	// MarkRevert denotes that the upgrade must be reverted. It will remove the upgrade from the character,
	// reduce the value of the trait and decrement its tier if able.
	MarkRevert = "-"

	// MarkSpecial denotes that the upgrade will benefit a special condition.
	// On application, the tier will not be touched. This mark is only used for characteristics.
	MarkSpecial = "*"
)

// marks holds the interpretable marks.
var marks = []string{
	MarkApply,
	MarkRevert,
	MarkSpecial,
}

// Upgrade describe one upgrade applied to the character. The mark indicate how
// to handle the upgrade considéring future upgrades. The cost is optionnal in
// certain cases.
type Upgrade struct {
	Mark string
	Name string
	Cost *int
	Line int
}

// parseUpgrade generate an upgrade from a raw line. The line must not be empty.
func parseUpgrade(line line) (Upgrade, error) {
	// Get the fields of the line
	fields := strings.Fields(line.Text)

	// The line shouldn't be empty
	if len(fields) == 0 {
		panic("empty line")
	}

	// The minimum number of fields is 2
	if len(fields) < 2 {
		return Upgrade{}, NewError(InvalidUpgradeFormat, line.Number)
	}

	// Parse the mark
	if !in(fields[0], marks) {
		return Upgrade{}, NewError(InvalidUpgradeMark, line.Number)
	}
	mark := fields[0]

	// Remove the field from the slice
	fields = fields[1:]

	// Check if a field seems to be a cost field
	var cost *int
	for i, field := range fields {
		// If one end has the brackets but not the other, that's an error:
		// brackets does by pairs, and are forbidden in the title
		if strings.HasPrefix(field, "[") != strings.HasSuffix(field, "]") {
			return Upgrade{}, NewError(InvalidUpgradeCost, line.Number)
		}

		// If the brackets are absents, that's not a cost, so skip the field.
		// Note that as both ends have brackets (or not), we just need to test
		// one of them.
		if !strings.HasPrefix(field, "[") {
			continue
		}

		// There can be only one cost on the line
		if cost != nil {
			return Upgrade{}, NewError(DuplicateUpgradeCost, line.Number)
		}

		// Check position of the cost
		if i != 0 && i != len(fields)-1 {
			return Upgrade{}, NewError(ForbidenCostPosition, line.Number)
		}

		// Trim the field to get the raw cost
		raw := strings.Trim(field, "[]")

		// Parse the cost
		c, err := strconv.Atoi(raw)
		if err != nil {
			return Upgrade{}, NewError(InvalidUpgradeCost, line.Number)
		}

		// Check the cost is positive
		if c < 0 {
			return Upgrade{}, NewError(InvalidUpgradeCost, line.Number)
		}
		cost = &c

		// Remove the field from the slice
		fields = append(fields[:i], fields[i+1:]...)
	}

	// The remaining line is the name of the upgrade
	if len(fields) == 0 {
		return Upgrade{}, NewError(EmptyUpgrade, line.Number)
	}

	// In case of non apply mark, the default cost value is 0.
	if mark != MarkApply && cost == nil {
		cost = IntP(0)
	}

	return Upgrade{
		Mark: mark,
		Name: strings.Join(fields, " "),
		Cost: cost,
		Line: line.Number,
	}, nil
}

// Split returns the name and speciality of an upgrade.
func (u Upgrade) Split() (string, string, error) {

	// Check if the skill has a speciality
	splits := strings.Split(u.Name, ":")
	if len(splits) > 2 {
		return "", "", NewError(InvalidUpgradeFormat, u.Line)
	}

	// Get name.
	name := splits[0]

	// Get speciality.
	var speciality string
	if len(splits) == 2 {
		speciality = splits[1]
	}

	return name, speciality, nil
}
