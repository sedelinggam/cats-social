package catsController

import (
	catHandler "cats-social/internal/delivery/http/v1/controller/cats/handler"
	catsService "cats-social/internal/service/cats"
	"cats-social/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func Init(group fiber.Router, catSvc catsService.CatService) {
	handler := catHandler.NewHandler(catSvc)
	user := group.Group("/cat")

	//Private Route
	jwt := auth.NewAuthMiddleware()
	privateRoute := user.Use(jwt)
	privateRoute.Post("/create", handler.CreateCat)
	privateRoute.Get("", handler.GetCats)

	//delete
	privateRoute.Delete("/delete/:id", handler.DeleteCat)
}
