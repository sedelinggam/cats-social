package v1

import (
	catsController "cats-social/internal/delivery/http/v1/controller/cats"
	matchesController "cats-social/internal/delivery/http/v1/controller/matches"
	usersController "cats-social/internal/delivery/http/v1/controller/users"
	catsService "cats-social/internal/service/cats"
	matchesService "cats-social/internal/service/matches"
	usersService "cats-social/internal/service/users"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(app *fiber.App, db *sqlx.DB) {
	var (
		usrSvc   = usersService.NewUserService(db)
		catSvc   = catsService.NewCatService(db)
		matchSvc = matchesService.NewMatchService(db)
	)
	v1 := app.Group("/v1")

	usersController.Init(v1, usrSvc)
	catsController.Init(v1, catSvc)
	matchesController.Init(v1, matchSvc)

}
