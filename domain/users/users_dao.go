package users

import (
	"fmt"

	usersdb "github.com/acargorkem/ecommerce_users-api/datasources/postgresql/users_db"
	"github.com/acargorkem/ecommerce_users-api/logger"
	"github.com/acargorkem/ecommerce_utils-go/rest_errors"

	postgresqlutils "github.com/acargorkem/ecommerce_users-api/utils/postgresql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users (first_name, last_name, email, created_at, hashed_password, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;"
	queryGetUser                = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE id=$1;"
	queryUpdateUser             = "UPDATE users SET first_name=$1, last_name=$2, email=$3, status=$4 WHERE id=$5;"
	queryDeleteUser             = "DELETE FROM users WHERE id=$1;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE status=$1 ORDER BY id ASC;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, created_at, hashed_password, status FROM users WHERE email=$1 AND status=$2;"
)

func (user *User) Get() *rest_errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Created_at, &user.Status); err != nil {
		logger.Error("error when trying to scan row on get user", err)
		return postgresqlutils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *rest_errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare insert user statement", err)
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.Created_at, user.Hashed_Password, user.Status)

	err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Hashed_Password, &user.Created_at)
	if err != nil {
		logger.Error("error when trying to scan row on save user", err)
		return postgresqlutils.ParseError(err)
	}

	return nil
}

func (user *User) Update() *rest_errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return postgresqlutils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *rest_errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return postgresqlutils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, postgresqlutils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Created_at, &user.Status); err != nil {
			logger.Error("error when trying to scan rows in find users by status", err)
			return nil, postgresqlutils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matchings status : %s ", status))
	}
	return results, nil
}

func (user *User) FindByEmail() *rest_errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare find user by email", err)
		return postgresqlutils.ParseError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Email, StatusActive)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email,
		&user.Created_at, &user.Hashed_Password, &user.Status); err != nil {
		logger.Error("error when trying to scan row on find user by email", err)
		return postgresqlutils.ParseError(err)
	}

	return nil
}
