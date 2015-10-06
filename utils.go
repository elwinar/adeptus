package main

// in checks whether a given string is in a given slice of strings.
func in(needle string, haystack []string) bool {
	for _, straw := range haystack {
		if needle == straw {
			return true
		}
	}
	return false
}
