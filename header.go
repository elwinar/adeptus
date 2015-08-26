package main

import (
	"fmt"
	"strings"
)

type Header struct {
	name       string
	origin     string
	background string
	role       string
}

// adds a metadata to the header
// raw: <pair#i><separator><value>
func (h Header) addMetadata(raw string) error {
}
