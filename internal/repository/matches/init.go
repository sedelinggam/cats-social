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
	GetMatchByCatID(ctx context.Context, catID string) (*entity.Match, error)
	GetMatchByID(ctx context.Context, matchID string) (*entity.Match, error)
	DeleteOtherMatches(ctx context.Context, userID string, receiverID string, matchID string) error
	Delete(ctx context.Context, id, userID string) error
	DeleteReceiverID(ctx context.Context, id, userID string) error
}

func NewRepository(db *sqlx.DB) MatchRepository {
	return &matchRepository{
		db: db,
	}
}
