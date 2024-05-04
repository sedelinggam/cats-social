package catsRepository

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type catRepository struct {
	db *sqlx.DB
}

type CatsRepository interface {
	Create(ctx context.Context, data entity.Cat) error
	Update(ctx context.Context, data entity.Cat) error
	Delete(ctx context.Context, catID, userID string) error
	GetById(ctx context.Context, id string) (*entity.Cat, error)
	GetCats(ctx context.Context, data request.GetCats) (*[]entity.Cat, error)
}

func NewRepository(db *sqlx.DB) CatsRepository {
	return &catRepository{
		db: db,
	}
}
