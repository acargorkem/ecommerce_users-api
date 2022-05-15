package users

import (
	usersdb "github.com/acargorkem/ecommerce_users-api/datasources/postgresql/users_db"
	dateutils "github.com/acargorkem/ecommerce_users-api/utils/date_utils"
	"github.com/acargorkem/ecommerce_users-api/utils/errors"
	postgresqlutils "github.com/acargorkem/ecommerce_users-api/utils/postgresql_utils"
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
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Created_at); err != nil {
		return postgresqlutils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	user.Created_at = dateutils.GetNowString()

	row := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.Created_at)

	err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Created_at)
	if err != nil {
		return postgresqlutils.ParseError(err)
	}

	return nil
}
