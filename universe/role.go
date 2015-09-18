package universe

import "adeptus/parser"

type Role struct {
	Name string
	Upgrades [][]parser.Upgrade
}

// NewRole returns the role associated to the name
func NewRole(name string) Role {
		return Role{}
}
