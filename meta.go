package main

import (
	"strings"
)

// Meta is a header and a collection of associated options.
type Meta struct {
	Label   string
	Options []string
	Line    int
}

// NewMeta returns a meta with name and options given the label.
func NewMeta(l line) (Meta, error) {

	if strings.Contains(l.Text, "(") {
		return Meta{}, NewError(InvalidBackgroundOptions, l.Number)
	}

	return Meta{
		Label: l.Text,
		Line:  l.Number,
	}, nil
}
