package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IdentifyCharacteristic returns the name, the value and the operand of a characteristic upgrade given it's label
func IdentifyCharacteristic(label string) (string, int, string, error) {

	// Check the characteristic contains a name and a value.
	splits := strings.Fields(label)
	if len(splits) != 2 {
		return "", 0, "", fmt.Errorf(`incorrect format for characteristic "%s": expecting name and value`, label)
	}

	// Retrieve the name
	name := splits[0]
	if name != strings.ToUpper(name) {
		return "", 0, "", fmt.Errorf(`characteristic "%s" name must be upper case`, label)
	}

	// Retrieve the value
	value, err := strconv.Atoi(splits[1])
	if err != nil {
		return "", 0, "", fmt.Errorf(`incorrect value for characteristic "%s"`, label)
	}

	// Retrive the operand
	var operand string
	switch {
	case strings.HasPrefix(splits[1], "+"):
		operand = "+"
	case strings.HasPrefix(splits[1], "-"):
		operand = "-"
	default:
		operand = "="
	}

	return name, value, operand, nil
}

// ApplyCharacteristicUpgrade returns the new value of origin after application an upgrade.
func ApplyCharacteristicUpgrade(before int, sign string, upgrade int) int {
	switch sign {
	case "+":
		before += upgrade
	case "-":
		before += upgrade
	case "=":
		before = upgrade
	default:
		panic(fmt.Sprintf("unrecognized sign %s", sign))
	}
	return before
}

// SplitUpgrade returns the name and speciality of an upgrade.
func SplitUpgrade(label string) (string, string, error) {

	// Check if the skill has a speciality
	splits := strings.Split(label, ":")
	if len(splits) > 2 {
		return "", "", fmt.Errorf(`incorrect format for upgrade "%s": expecting name or name: speciality`, label)
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

// countMatches return the number of matching aptitudes from two slices.
func countMatches(a []Aptitude, b []Aptitude) int {

	var m int
	for _, a := range a {
		for _, b := range b {
			if a == b {
				m++
			}
		}
	}
	return m
}
