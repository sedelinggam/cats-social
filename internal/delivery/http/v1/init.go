package v1

import (
	catsController "cats-social/internal/delivery/http/v1/controller/cats"
	usersController "cats-social/internal/delivery/http/v1/controller/users"
	catsService "cats-social/internal/service/cats"
	usersService "cats-social/internal/service/users"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(app *fiber.App, db *sqlx.DB) {
	var (
		usrSvc = usersService.NewUserService(db)
		catSvc = catsService.NewCatService(db)
	)
	v1 := app.Group("/v1")

	usersController.Init(v1, usrSvc)
	catsController.Init(v1, catSvc)
}
