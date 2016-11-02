package main

import "strings"

// comments holds the line separators describing a comment.
var comments = [2]string{
	"//",
	"#",
}

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

// Instruction returns the usefull (non-commented) part of the line.
func (l line) Instruction() string {
	for _, marker := range comments {
		pos := strings.Index(l.Text, marker)
		if pos == -1 {
			continue
		}

		if pos == 0 {
			return ""
		}

		l.Text = l.Text[:pos]
	}
	return l.Text
}

// IsEmpty check if the line contains no character or only blanks
func (l line) IsEmpty() bool {
	return len(strings.TrimSpace(l.Text)) == 0
}
