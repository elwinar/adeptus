package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
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
			Usage: "The dir location that contains the universe files.",
			Value: ".",
		},
	}

	app.Action = func(ctx *cli.Context) {
		_, c, err := Bootstrap(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Print()
	}

	app.Commands = []cli.Command{
		{
			Name:  "history",
			Usage: "display the history of a character sheet",
			Action: func(ctx *cli.Context) {
				_, c, err := Bootstrap(ctx)
				if err != nil {
					fmt.Println(err)
					return
				}
				c.PrintHistory()
			},
		},
		{
			Name:  "suggest",
			Usage: "display the list of purchasable upgrades, ordered by cost",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "max,m",
					Usage: "maximum XP cost of the upgrades to suggest",
				},
				cli.BoolFlag{
					Name:  "all,a",
					Usage: "display all upgrades regardless of their costs",
				},
				cli.BoolFlag{
					Name:  "with-spells,s",
					Usage: "display spells along with other upgrades",
				},
			},
			Action: func(ctx *cli.Context) {
				u, c, err := Bootstrap(ctx)
				if err != nil {
					fmt.Println(err)
					return
				}
				c.Suggest(u, ctx.Int("max"), ctx.Bool("all"), ctx.Bool("with-spells"))
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
