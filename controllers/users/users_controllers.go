package users

import (
	"net/http"

	"github.com/acargorkem/ecommerce_users-api/domain/users"
	"github.com/acargorkem/ecommerce_users-api/services"
	"github.com/acargorkem/ecommerce_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: HANDLE USER CREATION ERROR
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
