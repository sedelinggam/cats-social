package catsRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (cr catRepository) GetById(ctx context.Context, id string) (*entity.Cat, error) {
	var (
		resp entity.Cat
		err  error
	)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "id" = $1 AND "deleted_at" IS NULL`, resp.TableName())

	err = cr.db.GetContext(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
