package maildrop

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func fetchInbox(inbox string) Inbox {
	queryUrl := fmt.Sprintf("%s/mailbox/%s", baseurl, inbox)
	Logger.Println("fetchInbox:queryUrl:", queryUrl)

	body, status := doGetRequest(queryUrl)
	if status != http.StatusOK {
		Logger.Fatal("fetchInbox:", http.StatusText(status))
	}

	mInbox := Inbox{}
	err := json.Unmarshal(body, &mInbox)
	if err != nil {
		Logger.Fatal(err)
	}

	return mInbox
}
