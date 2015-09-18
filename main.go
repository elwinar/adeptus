package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"adeptus/parser"
)

func main() {
	app := cli.NewApp()

	app.Name = "Adeptus"
	app.Usage = "Compiles character sheet for Dark Heresy 2.0 related systems."
	app.Action = Display

	app.Commands = []cli.Command{
		{
			Name:   "history",
			Usage:  "Displays the history of the character.",
			Action: History,
			Flags: []cli.Flag{
				cli.StringFlag{
						Name: "character, c",
						Usage: "The filepath to the character sheet.",
				},
			},
		},
		{
			Name:   "display",
			Usage:  "Displays the current status of the character.",
			Action: Display,
			Flags: []cli.Flag{
				cli.StringFlag{
						Name: "character, c",
						Usage: "The filepath to the character sheet.",
				},
			},
		},
	}

	app.Run(os.Args)
}

// Compiles character sheet
func Display(ctx *cli.Context) {
	
	// Open and parse character sheet
	reader, err := os.Open(ctx.String("character"))
	if err != nil {
			log.Printf("undefined character: %s\n", err)
			return
	}
	
	sheet, err := parser.ParseSheet(reader)
	if err != nil {
			log.Printf("error character sheet: %s\n", err)
			return
	}
	
	// Create character with the seet
	character, err := NewCharacter(sheet)
	if err != nil {
			log.Printf("unable to create character: %s\n", err)
			return
	}
	
	character.Debug()
	
}

// Displays the history of the character
func History(ctx *cli.Context) {}
