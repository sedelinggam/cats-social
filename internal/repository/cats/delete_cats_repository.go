package catsRepository

import (
	"cats-social/internal/entity"
	"context"
	"errors"
	"fmt"
)

func (cr catRepository) Delete(ctx context.Context, catID, userID string) error {
	var (
		resp entity.Cat
		err  error
	)
	query := fmt.Sprintf(`UPDATE %s SET deleted_at = now() WHERE id = $1 AND user_id = $2`, resp.TableName())
	tx := cr.db.MustBegin()
	res, err := tx.ExecContext(ctx, query, catID, userID)
	if err != nil {
		return err
	}
	tx.Commit()

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("no rows in result set")
	}
	return nil
}
