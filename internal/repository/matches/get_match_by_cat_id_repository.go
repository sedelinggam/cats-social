package matchesRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (mr matchRepository) GetMatchByCatID(ctx context.Context, id string) (*entity.Match, error) {
	var (
		resp entity.Match
		err  error
	)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "user_cat_id" = $1`, resp.TableName())

	err = mr.db.GetContext(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
