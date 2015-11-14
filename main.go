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

	app.Before = Bootstrap

	app.Action = func(ctx *cli.Context) {
		character.Print()
	}

	app.Commands = []cli.Command{
		{
			Name:  "history",
			Usage: "display the history of a character sheet",
			Action: func(ctx *cli.Context) {
				character.PrintHistory()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Display character sheet.

// Display a character sheet history.
