package matchesRepository

import (
	"cats-social/internal/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type matchRepository struct {
	db *sqlx.DB
}

type MatchRepository interface {
	Create(ctx context.Context, data entity.Match) error
	GetMatches(ctx context.Context, userID string) ([]entity.Match, error)
}

func NewRepository(db *sqlx.DB) MatchRepository {
	return &matchRepository{
		db: db,
	}
}
