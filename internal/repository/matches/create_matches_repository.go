package matchesRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (cr matchRepository) Create(ctx context.Context, data entity.Match) error {
	query := fmt.Sprintf(`INSERT INTO %s(
		id,
		match_cat_id,
		user_cat_id,
		created_at,
		message
		) VALUES (
		:id,
		:match_cat_id,
		:user_cat_id,
		:created_at,
		:message
		)`, data.TableName())

	tx := cr.db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}
