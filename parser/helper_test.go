package parser

// IntP return the address of the given literal
func IntP(i int) *int {
	return &i
}

// StringP return the address of the given literal
func StringP(s string) *string {
	return &s
}
