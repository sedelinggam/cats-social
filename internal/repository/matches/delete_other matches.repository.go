package matchesRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (mr matchRepository) DeleteOtherMatches(ctx context.Context, userID string, receiverID string, matchID string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE (issuer_id = $1 OR receiver_id = $1) AND (issuer_id = $2 OR receiver_id = $2) AND id != $3`, entity.Match{}.TableName())
	tx := mr.db.MustBegin()
	_, err := tx.ExecContext(ctx, query, userID, receiverID, matchID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
