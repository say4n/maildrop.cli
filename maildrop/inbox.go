package maildrop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func fetchInbox(inbox string, logger *log.Logger) Inbox {
	queryUrl := fmt.Sprintf("%s/mailbox/%s", baseurl, inbox)
	logger.Println("fetchInbox:queryUrl:", queryUrl)

	req, err := createGetRequest(queryUrl)
	if err != nil {
		logger.Fatal(err)
	}

	client := getHTTPClient()
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
