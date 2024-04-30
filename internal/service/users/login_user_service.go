package usersService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/auth"
	"cats-social/pkg/lumen"
	"cats-social/pkg/password"
	"context"
	"errors"
)

func (us userService) Login(ctx context.Context, requestData request.UserLogin) (*response.UserAccessToken, error) {
	//Password Hash
	var (
		err error
	)

	// Find the user by credentials
	user, err := us.userRepo.GetUserByEmail(ctx, requestData.Email)
	if err != nil {
		//Duplicate unique key
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Compare password hash
	if !password.CheckPasswordHash(requestData.Password, user.Password) {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("password doesn't match"))
	}
	// Create the Claims
	accessToken, err := auth.GenerateToken(*user)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	return &response.UserAccessToken{
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: *accessToken,
	}, nil
}
