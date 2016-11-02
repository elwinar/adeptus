package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var (
	universe  Universe
	character *Character
)

// Bootstrap open and parse universe and character sheet.
func Bootstrap(ctx *cli.Context) error {
	// Open and parse the universe
	u, err := os.Open(ctx.GlobalString("universe"))
	if err != nil {
		return fmt.Errorf("%s %s", theme.Error("unable to open universe:"), err)
	}
	defer func() {
		_ = u.Close()
	}()
	universe, err = ParseUniverse(u)
	if err != nil {
		return fmt.Errorf("%s %s", theme.Error("corrupted universe:"), err)
	}

	// Open and parse character sheet.
	args := ctx.Args()
	if len(args) == 0 {
		return fmt.Errorf("%s undefined character", theme.Error("unable to open character sheet:"))
	}
	name := args[len(args)-1]
	c, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("%s %s", theme.Error("unable to open character sheet:"), err)
	}
	defer func() {
		_ = c.Close()
	}()
	sheet, err := ParseSheet(c)
	if err != nil {
		return fmt.Errorf("%s %s", theme.Error("corrupted character sheet:"), err)
	}

	// Create character with the sheet
	character, err = NewCharacter(universe, sheet)
	if err != nil {
		return fmt.Errorf("%s %s", theme.Error("unable to create character:"), err)
	}

	return nil
}
