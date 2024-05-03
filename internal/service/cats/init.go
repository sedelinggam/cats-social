package catsService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	catsRepository "cats-social/internal/repository/cats"
	matchesRepository "cats-social/internal/repository/matches"
	"context"

	"github.com/jmoiron/sqlx"
)

type catService struct {
	catRepo   catsRepository.CatsRepository
	matchRepo matchesRepository.MatchRepository
}

type CatService interface {
	CreateCat(ctx context.Context, requestData request.CreateCat) (*response.CreateCat, error)
	UpdateCat(ctx context.Context, requestData request.UpdateCat) (*response.UpdateCat, error)
	GetCats(ctx context.Context, requestData request.GetCats) (*[]response.GetCats, error)
}

func NewCatService(db *sqlx.DB) CatService {
	return &catService{
		catRepo:   catsRepository.NewRepository(db),
		matchRepo: matchesRepository.NewRepository(db),
	}
}
