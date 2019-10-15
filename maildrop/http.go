package maildrop

import (
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
		Timeout: time.Second * 2,
	}
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
