package postgresqlutils

import (
	"database/sql"
	"fmt"

	"github.com/acargorkem/ecommerce_users-api/utils/errors"
	"github.com/lib/pq"
)

const (
	UniqueViolationError = pq.ErrorCode("23505") // 'unique_violation'
)

func ParseError(err error) *errors.RestErr {
	if sqlErr, ok := err.(*pq.Error); ok {
		errMessage := fmt.Sprintf("%s : %s", sqlErr.Code.Name(), sqlErr.Detail)
		switch sqlErr.Code {
		case UniqueViolationError:
			return errors.NewBadRequestError(errMessage)
		}

		return errors.NewInternalServerError("database error")
	}
	if err == sql.ErrNoRows {
		return errors.NewNotFoundError("user not found")
	}
	return errors.NewInternalServerError("Whoops, something went wrong")
}
