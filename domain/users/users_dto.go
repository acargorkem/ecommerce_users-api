package users

import (
	"strings"

	"github.com/acargorkem/ecommerce_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id              int64  `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Hashed_Password string `json:"password"`
	Created_at      string `json:"created_at"`
	Status          string `json:"status"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
