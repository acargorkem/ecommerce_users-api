package postgresqlutils

import (
	"database/sql"
	"fmt"

	"github.com/acargorkem/ecommerce_utils-go/rest_errors"
	"github.com/lib/pq"
)

const (
	UniqueViolationError = pq.ErrorCode("23505") // 'unique_violation'
)

func ParseError(err error) *rest_errors.RestErr {
	if sqlErr, ok := err.(*pq.Error); ok {
		errMessage := fmt.Sprintf("%s : %s", sqlErr.Code.Name(), sqlErr.Detail)
		switch sqlErr.Code {
		case UniqueViolationError:
			return rest_errors.NewBadRequestError(errMessage)
		}

		return rest_errors.NewInternalServerError("errors on query", rest_errors.NewError("database_error"))
	}
	if err == sql.ErrNoRows {
		return rest_errors.NewNotFoundError("user not found")
	}
	return rest_errors.NewInternalServerError("unexpected_error", rest_errors.NewError("Whoops, something went wrong"))
}
