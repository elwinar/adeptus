package adeptus

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	
	app.Name = "Adeptus Manufactorum"
	app.Usage = "Compiles character sheet for Dark Heresy 2.0 related systems."
	app.Action = "Display"
	
	app.Commands = []cli.Command{
	{
		Name:  "history",
		Usage: "Displays the history of the character.",
		Action: History,
	}

	app.Run(os.Args)
}

// Compiles character sheet
func Display(c *cli.Context) {}

// Displays the history of the character
func History(c *cli.Context) {}