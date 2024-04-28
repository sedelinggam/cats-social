package usersHandler

import (
	usersService "cats-social/internal/service/users"
)

type userHandler struct {
	userService usersService.UserService
}

func NewHandler(usrSvc usersService.UserService) *userHandler {
	return &userHandler{
		userService: usrSvc,
	}
}
