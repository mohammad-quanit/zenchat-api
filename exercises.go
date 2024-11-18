package main

import (
	"errors"
	"fmt"

	// "net/http"
	"github.com/hypermodeinc/modus/sdk/go/pkg/http"
)

type Exercise struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Muscle       string `json:"muscle"`
	Equipment    string `json:"equipment"`
	Instructions string `json:"instructions"`
}

func GetExercises(muscle string) (*[]Exercise, error) {
	if muscle == "" {
		return nil, errors.New("muscle parameter is required")
	}

	exerciseEndpoint := "https://api.api-ninjas.com/v1/exercises?muscle=" + muscle
	req := http.NewRequest(exerciseEndpoint)

	res, err := http.Fetch(req)

	if err != nil {
		return nil, fmt.Errorf("unable to request on: %s due to %s", exerciseEndpoint, err.Error())
	}

	if !res.Ok() {
		return nil, fmt.Errorf("failed to fetch exercies. Received: %d %s", res.Status, res.StatusText)
	}

	// the API returns an array of exercises
	var exercises []Exercise
	res.JSON(&exercises)

	if len(exercises) == 0 {
		return nil, errors.New("expected at least one exerices in the response, but none were found")
	}

	return &exercises, nil
}
