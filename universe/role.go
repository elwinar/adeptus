package universe

import "adeptus/parser"

type Role struct {
	Name string
	Upgrades [][]parser.Upgrade
}
