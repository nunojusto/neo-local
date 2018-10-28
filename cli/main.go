package main

import (
	"log"
	"os"

	"github.com/CityOfZion/neo-local/cli/commands"
	"github.com/CityOfZion/neo-local/cli/logger"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

const (
	copyright   = "MIT"
	description = "Personal blockchain for NEO dApp development!"
)

var (
	// Version the tagged version of the binary.
	Version string

	author = cli.Author{
		Name: "City of Zion - https://github.com/cityofzion",
	}
	name = color.GreenString("neo-local")
)

func main() {
	logWriter := logger.NewWriter(name, Version)
	log.SetFlags(0)
	log.SetOutput(logWriter)

	app := cli.NewApp()
	app.Authors = []cli.Author{author}
	app.Commands = commands.GenerateCommandsIndex()
	app.Copyright = copyright
	app.Name = name
	app.Usage = description
	app.Version = Version

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(
			"%s %s. Please check the FAQ: https://github.com/CityOfZion/neo-local/wiki/FAQ",
			color.RedString("ERROR:"),
			color.RedString(err.Error()),
		)
	}
}
