package adeptus

// in checks whether a given string is in a given slice of strings
func in(needle string, haystack []string) bool {
	for _, straw := range haystack {
		if needle == straw {
			return true
		}
	}
	return false
}

// splitter return a function for the strings.FieldsFunc
func splitter(delimiters ...rune) func(c rune) bool {
	return func(c rune) bool {
		for _, d := range delimiters {
			if c == d {
				return true
			}
		}
		return false
	}
}
