package matchesService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	catRepository "cats-social/internal/repository/cats"
	matchesRepository "cats-social/internal/repository/matches"
	"context"

	"github.com/jmoiron/sqlx"
)

type matchService struct {
	matchRepo matchesRepository.MatchRepository
	catRepo   catRepository.CatsRepository
}

type MatchService interface {
	CreateMatch(ctx context.Context, requestData request.CreateMatch) (*response.CreateMatch, error)
	GetMatches(ctx context.Context, userID string) ([]response.GetMatches, error)
	DeleteMatch(ctx context.Context, matchID string) (*response.DeleteCat, error)
}

func NewMatchService(db *sqlx.DB) MatchService {
	return &matchService{
		matchRepo: matchesRepository.NewRepository(db),
		catRepo:   catRepository.NewRepository(db),
	}
}
