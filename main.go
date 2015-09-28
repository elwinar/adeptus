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

	// Open and parse character sheet.
	reader, err := os.Open(ctx.String("character"))
	if err != nil {
		log.Printf("cannot open character sheet: %s\n", err)
		return
	}

	sheet, err := ParseSheet(reader)
	if err != nil {
		log.Printf("unable to load character sheet: %s\n", err)
		return
	}

	// Open and parse the universe
	reader, err = os.Open("universe.json")
	if err != nil {
		log.Printf("cannot open  %s\n", err)
		return
	}

	universe, err := ParseUniverse(reader)
	if err != nil {
		log.Printf("unable to load  %s\n", err)
		return
	}

	// Create character with the sheet
	character, err := NewCharacter(universe, sheet)
	if err != nil {
		log.Printf("unable to create character: %s\n", err)
		return
	}

	character.Debug()

}
