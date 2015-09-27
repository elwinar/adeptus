package universe

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
