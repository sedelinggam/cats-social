package matchHandler

import (
	matchesService "cats-social/internal/service/matches"
)

type matchHandler struct {
	matchService matchesService.MatchService
}

func NewHandler(matchSvc matchesService.MatchService) *matchHandler {
	return &matchHandler{
		matchService: matchSvc,
	}
}
