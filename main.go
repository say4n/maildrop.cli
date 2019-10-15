package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/say4n/maildrop.cli/maildrop"
	"github.com/urfave/cli"
)

var (
	Version string
	Build   string
)

func main() {
	app := cli.NewApp()
	app.Name = "maildrop"
	app.Usage = "an unofficial cli client to maildrop.cc"
	app.Version = fmt.Sprintf("v%s (build %s)", Version, Build)

	logger := maildrop.GetLoggerInstance()

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
			Action:  maildrop.FetchInbox,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "address, a",
					Usage:    "mailbox address",
					Required: true,
				},
			},
		},
		{
			Name:    "view",
			Aliases: []string{"v"},
			Usage:   "read email from inbox",
			Action:  maildrop.FetchEmail,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "address, a",
					Usage:    "mailbox address",
					Required: true,
				},
				cli.StringFlag{
					Name:     "uid, u",
					Usage:    "unique ID of the email",
					Required: true,
				},
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete email from inbox",
			Action:  maildrop.DeleteEmail,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "address, a",
					Usage:    "mailbox address",
					Required: true,
				},
				cli.StringFlag{
					Name:     "uid, u",
					Usage:    "unique ID of the email",
					Required: true,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
