package catsRepository

import (
	"cats-social/internal/entity"
	"context"
	"errors"
	"fmt"
)

func (cr catRepository) Update(ctx context.Context, data entity.Cat) error {
	query := fmt.Sprintf(`UPDATE %s SET name = :name, race = :race, sex = :sex, age_in_month = :age_in_month, description = :description, image_urls = :image_urls WHERE id = :id`, data.TableName())

	tx := cr.db.MustBegin()
	res, err := tx.NamedExecContext(ctx, query, data)
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
