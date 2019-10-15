package maildrop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	au "github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

const (
	baseurl    = "https://api.maildrop.cc/v2"
	user_agent = "maildrop_cli"
	x_api_key  = "QM8VTHrLR2JloKTJMZ3N6Qa93FVsx8LapKCzEjui"
)

type Inbox struct {
	AltInbox string    `json:"altinbox"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Id      string `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}

func createGetRequest(url string) (*http.Request, error) {
	logger := GetLoggerInstance()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Fatal(err)
	}

	req.Header.Set("User-Agent", user_agent)
	req.Header.Set("x-api-key", x_api_key)

	return req, err
}

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

func fetchInbox(inbox string, logger *log.Logger) Inbox {
	queryUrl := fmt.Sprintf("%s/%s/%s", baseurl, "mailbox", inbox)
	logger.Println("fetchInbox:queryUrl:", queryUrl)

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := createGetRequest(queryUrl)
	if err != nil {
		logger.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		logger.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Fatal(err)
	}

	mInbox := Inbox{}
	err = json.Unmarshal(body, &mInbox)
	if err != nil {
		logger.Fatal(err)
	}

	return mInbox
}

func FetchEmail(c *cli.Context) error {
	logger := GetLoggerInstance()
	if len(c.Args()) < 2 {
		logger.Fatal("FetchEmail.error: cannot show email without inbox name and email uid")
	}
	logger.Printf("inbox: %s, email: %s", c.Args()[0], c.Args()[1])

	return nil
}
