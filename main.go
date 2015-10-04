package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "adeptus"
	app.Usage = "Compiles character sheet for Dark Heresy 2.0 related systems."
	app.Version = "alpha"
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
		log.Println(err)
		return
	}
}

// Display character sheet.
func Display(ctx *cli.Context) {

	// Open and parse the universe
	u, err := os.Open(ctx.GlobalString("universe"))
	if err != nil {
		log.Println("error:", "unable to open universe:", err)
		return
	}
	defer func() {
		_ = u.Close()
	}()
	universe, err := ParseUniverse(u)
	if err != nil {
		log.Println("error:", "corrupted universe:", err)
		return
	}

	// Open and parse character sheet.
	c, err := os.Open(ctx.String("character"))
	if err != nil {
		log.Println("error:", "unable to open character sheet:", err)
		return
	}
	defer func() {
		_ = c.Close()
	}()
	sheet, err := ParseSheet(c)
	if err != nil {
		log.Println("error:", "corrupted character sheet:", err)
		return
	}

	// Create character with the sheet
	character, err := NewCharacter(universe, sheet)
	if err != nil {
		log.Println("error:", "unable to create character:", err)
		return
	}

	// Print the character sheet on screen
	character.Print()
}
