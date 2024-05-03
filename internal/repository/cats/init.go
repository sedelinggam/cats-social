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
	Update(ctx context.Context, data entity.Cat) error
}

func NewRepository(db *sqlx.DB) CatsRepository {
	return &catRepository{
		db: db,
	}
}
