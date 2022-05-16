package users

import (
	usersdb "github.com/acargorkem/ecommerce_users-api/datasources/postgresql/users_db"
	dateutils "github.com/acargorkem/ecommerce_users-api/utils/date_utils"
	"github.com/acargorkem/ecommerce_users-api/utils/errors"
	postgresqlutils "github.com/acargorkem/ecommerce_users-api/utils/postgresql_utils"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, created_at) VALUES ($1, $2, $3, $4) RETURNING *;"
	queryGetUser    = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=$1;"
	queryUpdateUser = "UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4;"
	queryDeleteUser = "DELETE FROM users WHERE id=$1;"
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

func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return postgresqlutils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return postgresqlutils.ParseError(err)
	}
	return nil
}
