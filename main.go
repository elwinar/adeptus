package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "adeptus"
	app.Usage = "Compiles character sheet for Dark Heresy 2.0 related systems."
	app.Version = "beta"
	app.Authors = []cli.Author{
		{
			Name:  "Romain Baugue",
			Email: "romain.baugue@gmail.com",
		},
		{
			Name:  "Alexandre Thomas",
			Email: "alexandre.thomas@outlook.fr",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "character, c",
			Usage: "The filepath to the character sheet.",
		},
		cli.StringFlag{
			Name:  "universe, u",
			Usage: "The filepath to the character universe.",
			Value: "universe.json",
		},
	}
	app.Action = Display

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Display character sheet.
func Display(ctx *cli.Context) {

	// Open and parse the universe
	u, err := os.Open(ctx.GlobalString("universe"))
	if err != nil {
		fmt.Println(theme.Error("unable to open universe:"), err)
		os.Exit(1)
	}
	defer func() {
		_ = u.Close()
	}()
	universe, err := ParseUniverse(u)
	if err != nil {
		fmt.Println(theme.Error("corrupted universe:"), err)
		os.Exit(1)
	}

	// Open and parse character sheet.
	c, err := os.Open(ctx.String("character"))
	if err != nil {
		fmt.Println(theme.Error("unable to open character sheet:"), err)
		os.Exit(1)
	}
	defer func() {
		_ = c.Close()
	}()
	sheet, err := ParseSheet(c)
	if err != nil {
		fmt.Println(theme.Error("corrupted character sheet:"), err)
		os.Exit(1)
	}

	// Create character with the sheet
	character, err := NewCharacter(universe, sheet)
	if err != nil {
		fmt.Println(theme.Error("unable to create character:"), err)
		os.Exit(1)
	}

	// Print the character sheet on screen
	character.Print()
}
