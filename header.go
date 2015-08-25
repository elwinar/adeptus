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
// raw: <pair>:<value>
func (h Header) addMetadata(raw string) error {
	pair := strings.SplitN(raw, ":", 2)
	if len(pair) < 2 {
		err := fmt.Errorf("Error in parsing of header. Expected pair key:value, having pair without value")
		return err
	}
	key := pair[0]
	value := strings.TrimSpace(pair[1])
	switch key {
	case "name":
		h.name = value
	case "origin":
		h.origin = value
	case "background":
		h.background = value
	case "role":
		h.role = value
	default:
		err := fmt.Errorf("Undefined key: %s in header.", key)
		return err
	}
	return nil
}
