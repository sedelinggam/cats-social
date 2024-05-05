package usersRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (ur userRepository) Create(ctx context.Context, data entity.User) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, email, name, password, created_at) VALUES (:id, :email, :name, :password, :created_at)`, data.TableName())

	tx := ur.db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
