package users

import (
	"fmt"

	usersdb "github.com/acargorkem/ecommerce_users-api/datasources/postgresql/users_db"
	"github.com/acargorkem/ecommerce_users-api/utils/errors"
	postgresqlutils "github.com/acargorkem/ecommerce_users-api/utils/postgresql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users (first_name, last_name, email, created_at, status) VALUES ($1, $2, $3, $4, $5) RETURNING *;"
	queryGetUser          = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE id=$1;"
	queryUpdateUser       = "UPDATE users SET first_name=$1, last_name=$2, email=$3, status=$4 WHERE id=$5;"
	queryDeleteUser       = "DELETE FROM users WHERE id=$1;"
	queryFindUserbyStatus = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE status=$1 ORDER BY id ASC;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Created_at, &user.Status); err != nil {
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

	row := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.Created_at, user.Status)

	err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Created_at)
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

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)
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

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserbyStatus)
	if err != nil {
		return nil, postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, postgresqlutils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Created_at, &user.Status); err != nil {
			return nil, postgresqlutils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matchings status : %s ", status))
	}
	return results, nil

}
