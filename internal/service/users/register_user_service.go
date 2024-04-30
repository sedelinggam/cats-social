package usersService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/internal/entity"
	"cats-social/pkg/auth"
	"cats-social/pkg/lumen"
	"cats-social/pkg/password"
	"context"
	"time"

	"github.com/google/uuid"
)

func (us userService) Register(ctx context.Context, requestData request.UserRegister) (*response.UserAccessToken, error) {
	var (
		err          error
		hashPassword string
	)
	//Password Hash
	hashPassword, err = password.HashPassword(requestData.Password)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	//Create User
	userData := entity.User{
		ID:        uuid.New().String(),
		Email:     requestData.Email,
		Name:      requestData.Name,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}
	err = us.userRepo.Create(ctx, userData)
	if err != nil {
		//Duplicate unique key
		if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	// Create the Claims
	accessToken, err := auth.GenerateToken(userData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	return &response.UserAccessToken{
		Email:       requestData.Email,
		Name:        requestData.Name,
		AccessToken: *accessToken,
	}, nil
}
