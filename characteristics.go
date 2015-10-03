package main

import (
	"strconv"
	"strings"
)

// Characteristics is the block of characteristic upgrades after the header
type Characteristics []Upgrade

// ParseHeader generate a Characteristics from a block of lines. The block must not be
// empty.
func parseCharacteristics(block []line) (Characteristics, error) {
	// Check the block is non-empty
	if len(block) == 0 {
		panic("empty block")
	}

	// Parse each upgrade
	upgrades := []Upgrade{}
	for _, line := range block {

		// The line should be made of label and value
		splits := strings.Fields(line.Text)
		if len(splits) != 2 {
			return Characteristics{}, NewError(InvalidCharacteristicFormat, line.Number)
		}

		// Check the value is numeric
		if strings.ContainsAny(splits[1], "+|-") {
			return Characteristics{}, NewError(NotIntegerCharacteristicValue, line.Number)
		}
		_, err := strconv.Atoi(splits[1])
		if err != nil {
			return Characteristics{}, NewError(NotIntegerCharacteristicValue, line.Number)
		}

		u := Upgrade{
			Mark: "-",
			Name: strings.Join(splits, " "),
			Cost: nil,
			Line: line.Number,
		}

		// Add the characteristic to the list
		upgrades = append(upgrades, u)
	}

	return upgrades, nil
}
