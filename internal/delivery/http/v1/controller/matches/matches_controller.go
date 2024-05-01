package matchesController

import (
	matchHandler "cats-social/internal/delivery/http/v1/controller/matches/handler"
	matchesService "cats-social/internal/service/matches"
	"cats-social/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func Init(group fiber.Router, matchSvc matchesService.MatchService) {
	handler := matchHandler.NewHandler(matchSvc)
	user := group.Group("/cat/match")

	//Private Route
	jwt := auth.NewAuthMiddleware()
	privateRoute := user.Use(jwt)
	privateRoute.Post("/", handler.CreateMatch)
	privateRoute.Get("/", handler.GetMatches)
}
