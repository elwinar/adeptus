package universe

import "adeptus/parser"

// Role represents a character's role
type Role struct {
	Name     string
	Upgrades [][]parser.Upgrade
}

// NewRole returns the role associated to the name
func NewRole(name string) Role {
	return Role{
		Name: name,
	}
}
