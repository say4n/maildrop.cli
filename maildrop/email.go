package maildrop

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

type Email struct {
	Id         string `json:"id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Subject    string `json:"subject"`
	Date       string `json:"date"`
	Body       string `json:"body"`
	Html       string `json:"html"`
	Plaintext  string
	ParsedDate string
}

func fetchEmail(inbox, email_uid string) Email {
	queryUrl := fmt.Sprintf("%s/mailbox/%s/%s", baseurl, inbox, email_uid)
	Logger.Println("fetchEmail:queryUrl:", queryUrl)

	body, status := doGetRequest(queryUrl)
	if status != http.StatusOK {
		Logger.Fatal("fetchEmail:", http.StatusText(status))
	}

	mEmail := Email{}
	err := json.Unmarshal(body, &mEmail)
	if err != nil {
		Logger.Fatal(err)
	}

	p := bluemonday.StrictPolicy()
	mEmail.Plaintext = html.UnescapeString(p.Sanitize(mEmail.Html))

	format := "2006-01-02T15:04:05Z"
	t, err := time.Parse(format, mEmail.Date)
	if err != nil {
		fmt.Println(err)
	}

	mEmail.ParsedDate = t.Format(time.RFC1123)

	return mEmail
}
