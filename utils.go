package main

import "strings"

// in checks whether a given string is in a given slice of strings.
func in(needle string, haystack []string) bool {
	for _, straw := range haystack {
		if needle == straw {
			return true
		}
	}
	return false
}

func split(s string, c rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == c
	})
}

// IntP returns the pointer to the given var
func IntP(v int) *int {
	return &v
}
