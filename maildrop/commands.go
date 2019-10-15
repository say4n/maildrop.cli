package maildrop

import (
	"fmt"
	_ "log"
	"os"

	au "github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

const baseurl = "https://api.maildrop.cc/v2"

func FetchInbox(c *cli.Context) error {
	if len(c.Args()) == 0 {
		Logger.Fatal("FetchInbox.error: cannot list emails without inbox name")
	}

	mInbox := fetchInbox(c.Args().First())

	fmt.Printf("Alias Address: %s@maildrop.cc\n", au.Bold(mInbox.AltInbox))
	fmt.Printf("Email(s) for %s@maildrop.cc:\n", au.BrightBlue(c.Args().First()))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Email UID", "Subject"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	for _, msg := range mInbox.Messages {
		uid := fmt.Sprintf("%s", au.Bold(msg.Id))
		text := fmt.Sprintf("%s by %s", msg.Subject, msg.From)
		table.Append([]string{uid, text})
	}

	table.Render()

	return nil
}

func FetchEmail(c *cli.Context) error {
	if len(c.Args()) < 2 {
		Logger.Fatal("FetchEmail.error: cannot show email without inbox name and email uid")
	}

	inbox := c.Args()[0]
	email_uid := c.Args()[1]
	Logger.Printf("inbox: %s, email: %s", inbox, email_uid)

	mEmail := fetchEmail(inbox, email_uid)

	fmt.Printf("%s: %s\n", au.Bold("subject"), mEmail.Subject)
	fmt.Printf("%s: %s\n", au.Bold("from"), mEmail.From)
	fmt.Printf("%s: %s\n", au.Bold("date"), mEmail.ParsedDate)
	fmt.Printf("%s:\n%s\n", au.Bold("message"), mEmail.Plaintext)

	return nil
}
