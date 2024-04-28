package usersController

import (
	usersHandler "cats-social/internal/delivery/http/v1/controller/users/handler"
	usersService "cats-social/internal/service/users"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(group fiber.Router) {
	handler := usersHandler.NewHandler(usersService.NewUserService(&sqlx.DB{}))
	user := group.Group("/user")

	publicRoute := user
	publicRoute.Post("/register", handler.Register)
	publicRoute.Post("/login", handler.Register)
}
