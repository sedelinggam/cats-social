package catsRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (cr catRepository) Create(ctx context.Context, data entity.Cat) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, user_id, name, race, sex, age_in_month, description, is_already_matched, image_urls, created_at) VALUES (:id, :user_id, :name, :race, :sex, :age_in_month, :description, :is_already_matched, :image_urls, :created_at)`, data.TableName())

	_, err := cr.db.NamedQueryContext(ctx, query, data)
	if err != nil {
		return err
	}

	return nil
}
