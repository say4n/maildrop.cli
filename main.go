package main

import (
	"log"
	"os"

	"github.com/say4n/maildrop.cli/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "maildrop"
	app.Usage = "an unofficial cli client to maildrop.cc"

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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
