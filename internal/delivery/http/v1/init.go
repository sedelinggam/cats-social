package v1

import (
	usersController "cats-social/internal/delivery/http/v1/controller/users"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	v1 := app.Group("/v1")

	usersController.Init(v1)
}
