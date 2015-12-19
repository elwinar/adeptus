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
		{
			Name:  "suggest",
			Usage: "display the list of purchasable upgrades, ordered by cost",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "max",
					Usage: "maximum XP cost of the upgrades to suggest",
				},
				cli.BoolFlag{
					Name:  "all",
					Usage: "display all upgrades regardless of their costs",
				},
			},
			Action: func(ctx *cli.Context) {
				character.Suggest(ctx.Int("max"), ctx.Bool("all"))
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
