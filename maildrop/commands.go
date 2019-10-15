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
	address := c.String("address")
	mInbox := fetchInbox(address)

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
	inbox := c.String("address")
	email_uid := c.String("uid")
	Logger.Printf("inbox: %s, email: %s", inbox, email_uid)

	mEmail := fetchEmail(inbox, email_uid)

	fmt.Printf("%s: %s\n", au.Bold("subject"), mEmail.Subject)
	fmt.Printf("%s: %s\n", au.Bold("from"), mEmail.From)
	fmt.Printf("%s: %s\n", au.Bold("date"), mEmail.ParsedDate)
	fmt.Printf("%s:\n%s\n", au.Bold("message"), mEmail.Plaintext)

	return nil
}

func DeleteEmail(c *cli.Context) error {
	inbox := c.String("address")
	email_uid := c.String("uid")
	Logger.Printf("inbox: %s, email: %s", inbox, email_uid)

	deleteEmail(inbox, email_uid)

	return nil
}
