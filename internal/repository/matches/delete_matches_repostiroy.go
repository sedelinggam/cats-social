package matchesRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (mr matchRepository) Delete(ctx context.Context, matchID, userID string) error {
	var (
		resp entity.Match
		err  error
	)
	query := fmt.Sprintf(`DELETE FROM %s where id = $1 AND issuer_id = $id`, resp.TableName())

	tx := mr.db.MustBegin()
	_, err = tx.ExecContext(ctx, query, matchID, userID)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}
