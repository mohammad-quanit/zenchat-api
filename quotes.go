package main

import (
	"errors"
	"fmt"

	"github.com/hypermodeinc/modus/sdk/go/pkg/http"
)

type Quote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

// this function makes a request to an API that returns data in JSON format, and
// returns an object representing the data
func GetRandomQuote() (*Quote, error) {
	req := http.NewRequest("https://zenquotes.io/api/random")

	res, err := http.Fetch(req)
	if err != nil {
		return nil, err
	}

	if !res.Ok() {
		return nil, fmt.Errorf("failed to fetch quote. Received: %d %s", res.Status, res.StatusText)
	}

	// the API returns an array of quotes, but we only want the first one
	var quotes []Quote
	res.JSON(&quotes)

	if len(quotes) == 0 {
		return nil, errors.New("expected at least one quote in the response, but none were found")
	}

	return &quotes[0], nil
}
