package maildrop

import (
	"fmt"
	_ "log"

	au "github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

const baseurl = "https://api.maildrop.cc/v2"

func FetchInbox(c *cli.Context) error {
	logger := GetLoggerInstance()
	if len(c.Args()) == 0 {
		logger.Fatal("FetchInbox.error: cannot list emails without inbox name")
	}

	mInbox := fetchInbox(c.Args().First(), logger)

	fmt.Printf("Alias Address: %s@maildrop.cc\n", au.Bold(mInbox.AltInbox))
	fmt.Printf("Email(s) for %s@maildrop.cc:\n", au.BrightBlue(c.Args().First()))

	for _, msg := range mInbox.Messages {
		fmt.Printf("[%s]\t%s by %s\n", au.Bold(msg.Id), msg.Subject, msg.From)
	}

	return nil
}

func FetchEmail(c *cli.Context) error {
	logger := GetLoggerInstance()
	if len(c.Args()) < 2 {
		logger.Fatal("FetchEmail.error: cannot show email without inbox name and email uid")
	}
	logger.Printf("inbox: %s, email: %s", c.Args()[0], c.Args()[1])

	return nil
}
