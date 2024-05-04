package matchesRepository

import (
	"cats-social/internal/entity"
	"context"
	"errors"
	"fmt"
)

func (mr matchRepository) DeleteReceiverID(ctx context.Context, matchID, userID string) error {
	var (
		resp entity.Match
		err  error
	)
	query := fmt.Sprintf(`DELETE FROM %s where id = $1 AND receiver_id = $2`, resp.TableName())

	tx := mr.db.MustBegin()
	res, err := tx.ExecContext(ctx, query, matchID, userID)
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
