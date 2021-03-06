package services

import (
	"github.com/acargorkem/ecommerce_users-api/domain/users"
	cryptoutils "github.com/acargorkem/ecommerce_users-api/utils/crypto_utils"
	dateutils "github.com/acargorkem/ecommerce_users-api/utils/date_utils"
	"github.com/acargorkem/ecommerce_utils-go/rest_errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *rest_errors.RestErr)
	GetUser(int64) (*users.User, *rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *rest_errors.RestErr)
	DeleteUser(int64) *rest_errors.RestErr
	SearchUser(string) (users.Users, *rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *rest_errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err

	}

	hashedPassword, err := cryptoutils.HashPassword(user.Hashed_Password)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("errors when trying to hash password", rest_errors.NewError("Whoops, something went wrong"))
	}

	user.Status = users.StatusActive
	user.Created_at = dateutils.GetNowDbFormat()
	user.Hashed_Password = hashedPassword

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *rest_errors.RestErr) {
	result := &users.User{
		Id: userId,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *rest_errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *rest_errors.RestErr) {
	dao := &users.User{
		Email: request.Email,
	}
	if err := dao.FindByEmail(); err != nil {
		return nil, err
	}

	isInvalidPassword := cryptoutils.CheckPassword(request.Password, dao.Hashed_Password)
	if isInvalidPassword != nil {
		return nil, rest_errors.NewUnauthorizedError("invalid credentials")
	}

	return dao, nil
}
