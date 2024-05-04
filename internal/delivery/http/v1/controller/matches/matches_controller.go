package matchesController

import (
	matchHandler "cats-social/internal/delivery/http/v1/controller/matches/handler"
	matchesService "cats-social/internal/service/matches"
	"cats-social/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func Init(group fiber.Router, matchSvc matchesService.MatchService) {
	handler := matchHandler.NewHandler(matchSvc)
	match := group.Group("/cat/match")

	//Private Route
	jwt := auth.NewAuthMiddleware()
	privateRoute := match.Use(jwt)
	privateRoute.Post("/", handler.CreateMatch)
	privateRoute.Get("/", handler.GetMatches)
	privateRoute.Post("/approve", handler.ApproveMatch)
}
