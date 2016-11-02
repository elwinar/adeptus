package main

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
