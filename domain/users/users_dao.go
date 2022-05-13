package users

import (
	"fmt"
	"strings"

	usersdb "github.com/acargorkem/ecommerce_users-api/datasources/postgresql/users_db"
	dateutils "github.com/acargorkem/ecommerce_users-api/utils/date_utils"
	"github.com/acargorkem/ecommerce_users-api/utils/errors"
)

const (
	indexUniqueEmail = "duplicate key value violates unique constraint \"users_email_key\""
	queryInsertUser  = "INSERT INTO users (first_name, last_name, email, created_at) VALUES ($1, $2, $3, $4) RETURNING *;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.Created_at = result.Created_at
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.Created_at = dateutils.GetNowString()

	row := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.Created_at)

	err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Created_at)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get last insert user: %s", err.Error()))
	}

	return nil
}
