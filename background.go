package main

// Background represents an element providing traits to a character.
type Background struct {
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Upgrades []string `json:"upgrades"`
}
