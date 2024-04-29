package usersController

import (
	usersHandler "cats-social/internal/delivery/http/v1/controller/users/handler"
	usersService "cats-social/internal/service/users"

	"github.com/gofiber/fiber/v2"
)

func Init(group fiber.Router, usrSvc usersService.UserService) {
	handler := usersHandler.NewHandler(usrSvc)
	user := group.Group("/user")

	publicRoute := user
	publicRoute.Post("/register", handler.Register)
	publicRoute.Post("/login", handler.Login)
}
