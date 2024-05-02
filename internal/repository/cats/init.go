package catsRepository

import (
	"cats-social/internal/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type catRepository struct {
	db *sqlx.DB
}

type CatsRepository interface {
	Create(ctx context.Context, data entity.Cat) error
	Delete(ctx context.Context, catID string) error
}

func NewRepository(db *sqlx.DB) CatsRepository {
	return &catRepository{
		db: db,
	}
}
