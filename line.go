package main

import "strings"

type line struct {
	Number int
	Text   string
}

// newLine return a wrapper for the line string
func newLine(text string, number int) line {
	return line{
		Text:   text,
		Number: number,
	}
}

// IsComment check that the line starts with a comment indicator
func (l line) IsComment() bool {
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

// IsEmpty check if the line contains no character or only blanks
func (l line) IsEmpty() bool {
	return len(strings.TrimSpace(l.Text)) == 0
}
