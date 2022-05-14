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
	errorNorows      = "no rows in result"
	queryInsertUser  = "INSERT INTO users (first_name, last_name, email, created_at) VALUES ($1, $2, $3, $4) RETURNING *;"
	queryGetUser     = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=$1;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Created_at); err != nil {
		if strings.Contains(err.Error(), errorNorows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

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
