package main

import (
	"errors"
	"fmt"

	"github.com/hypermodeinc/modus/sdk/go/pkg/http"
	// "github.com/hypermodeinc/modus/sdk/go/pkg/postgresql"
)

type Person struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     int    `json:"email"`
	Gender    string `json:"gender"`
}

func GetPerson(name string) (*Person, error) {
	endpoint := fmt.Sprintf("https://lwyobyydvdriwaukdbjf.supabase.co/rest/v1/users?username=eq.%sselect=*", name)
	fmt.Println("endpoint...", endpoint)
	req := http.NewRequest(endpoint)

	res, err := http.Fetch(req)
	if err != nil {
		return nil, err
	}

	if !res.Ok() {
		return nil, fmt.Errorf("failed to fetch user. Received: %d %s", res.Status, res.StatusText)
	}

	// the API returns an array of quotes, but we only want the first one
	var person []Person
	res.JSON(&person)

	if len(person) == 0 {
		return nil, errors.New("expected at least one quote in the response, but none were found")
	}

	return &person[0], nil
}
