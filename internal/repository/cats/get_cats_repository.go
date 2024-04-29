package catsRepository

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (cr catRepository) GetCats(ctx context.Context, data request.GetCats) (*[]entity.Cat, error) {
	var (
		resp []entity.Cat
		err  error
	)
	query := fmt.Sprintf(`SELECT * FROM cats`)

	err = cr.db.GetContext(ctx, &resp, query, data)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
