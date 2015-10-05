package main

// IntP return the address of the given literal
func IntP(i int) *int {
	return &i
}

// StringP return the address of the given literal
func StringP(s string) *string {
	return &s
}

func newMeta(text string) Meta {
	m, err := NewMeta(newLine(text, 0))
	if err != nil {
		panic(err)
	}
	return m
}
