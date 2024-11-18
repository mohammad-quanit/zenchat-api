package main

import (
	"errors"
	// "fmt"
	// "github.com/google/uuid"
	"github.com/hypermodeinc/modus/sdk/go/pkg/postgresql"
)

type User struct {
	// Id        uuid.UUID `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

const connection = "db-connection"

func GetUser(id string) (*User, error) {
	const query = "select * from users where id = $1"
	rows, _, err := postgresql.Query[User](connection, query, id)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	var user = rows[0]
	return &user, nil
}
