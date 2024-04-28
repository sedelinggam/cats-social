package usersService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"context"
)

func (us userService) Login(ctx context.Context, requestData request.UserLogin) (*response.UserAccessToken, error) {
	return nil, nil
}
