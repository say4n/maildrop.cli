package main

import (
	"io/ioutil"
	"os"

	"github.com/say4n/maildrop.cli/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "maildrop"
	app.Usage = "an unofficial cli client to maildrop.cc"

	logger := commands.GetLoggerInstance()

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "logging, l",
			Usage: "enable logging",
		},
	}

	app.Before = func(c *cli.Context) error {
		if !c.Bool("logging") {
			logger.SetFlags(0)
			logger.SetOutput(ioutil.Discard)
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "inbox",
			Aliases: []string{"i"},
			Usage:   "show emails in inbox",
			Action:  commands.FetchInbox,
		},
		{
			Name:    "view",
			Aliases: []string{"v"},
			Usage:   "list emails from inbox",
			Action:  commands.FetchEmail,
		},
	}

	// if

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
