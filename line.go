package adeptus

import "strings"

type Line struct {
	Number int
	Text   string
}

// Checks if the line is a comment
func (l Line) IsComment() bool {
	if strings.HasPrefix(l.Text, "//") {
		return true
	}
	if strings.HasPrefix(l.Text, "#") {
		return true
	}
	if strings.HasPrefix(l.Text, ";") {
		return true
	}
	return false
}

// Checks if the line is empty
func (l Line) IsEmpty() bool {
	return len(l.Text) == 0
}
