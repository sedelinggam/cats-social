package catsService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	catsRepository "cats-social/internal/repository/cats"
	"context"

	"github.com/jmoiron/sqlx"
)

type catService struct {
	catRepo catsRepository.CatsRepository
}

type CatService interface {
	CreateCat(ctx context.Context, requestData request.CreateCat) (*response.CreateCat, error)
	DeleteCat(ctx context.Context, catID string) (*response.DeleteCat, error)
}

func NewCatService(db *sqlx.DB) CatService {
	return &catService{
		catRepo: catsRepository.NewRepository(db),
	}
}
