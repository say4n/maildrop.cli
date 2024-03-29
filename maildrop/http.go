package maildrop

import (
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var httpClient *http.Client
var honce sync.Once

const (
	user_agent = "maildrop_cli"
	x_api_key  = "QM8VTHrLR2JloKTJMZ3N6Qa93FVsx8LapKCzEjui"
)

func getHTTPClient() *http.Client {
	honce.Do(func() {
		httpClient = createHTTPClient()
	})

	return httpClient
}

func createHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5,
	}
}

func createGetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		Logger.Fatal(err)
	}

	req.Header.Set("User-Agent", user_agent)
	req.Header.Set("x-api-key", x_api_key)

	return req, err
}

func createDeleteRequest(url, referer string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		Logger.Fatal(err)
	}

	req.Header.Set("User-Agent", user_agent)
	req.Header.Set("x-api-key", x_api_key)
	req.Header.Set("Referer", referer)

	return req, err
}

func doGetRequest(url string) ([]byte, int) {
	req, err := createGetRequest(url)
	if err != nil {
		Logger.Fatal(err)
	}

	client := getHTTPClient()
	res, err := client.Do(req)
	if err != nil {
		Logger.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Logger.Fatal(err)
	}

	return body, res.StatusCode
}

func doDeleteRequest(url, referer string) int {
	req, err := createDeleteRequest(url, referer)
	if err != nil {
		Logger.Fatal(err)
	}

	client := getHTTPClient()
	res, err := client.Do(req)
	if err != nil {
		Logger.Fatal(err)
	}

	return res.StatusCode
}
