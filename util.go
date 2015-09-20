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
