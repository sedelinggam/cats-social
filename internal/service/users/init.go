package usersService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	usersRepository "cats-social/internal/repository/users"
	"context"

	"github.com/jmoiron/sqlx"
)

type userService struct {
	userRepo usersRepository.UserRepository
}

type UserService interface {
	Login(ctx context.Context, requestData request.UserLogin) (*response.UserAccessToken, error)
	Register(ctx context.Context, requestData request.UserRegister) (*response.UserAccessToken, error)
}

func NewUserService(db *sqlx.DB) UserService {
	return &userService{
		userRepo: usersRepository.NewRepository(db),
	}
}
