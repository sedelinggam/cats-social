package usersRepository

import (
	"cats-social/internal/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	Create(ctx context.Context, data entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

func NewRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
