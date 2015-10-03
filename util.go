package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IdentifyCharacteristic returns the name, the value and the operand of a characteristic upgrade given it's label
func IdentifyCharacteristic(u Upgrade) (string, int, string, error) {

	// Check the characteristic contains a name and a value.
	splits := strings.Fields(u.Name)
	if len(splits) != 2 {
		return "", 0, "", NewError(InvalidCharacteristicFormat, u.Line)
	}

	// Retrieve the name
	name := splits[0]
	if name != strings.ToUpper(name) {
		return "", 0, "", NewError(InvalidCharacteristicCase, u.Line)
	}

	// Retrieve the value
	value, err := strconv.Atoi(splits[1])
	if err != nil {
		return "", 0, "", NewError(InvalidCharacteristicValue, u.Line)
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
