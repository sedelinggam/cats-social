package catsRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (cr catRepository) Delete(ctx context.Context, catID string) error {
    query := fmt.Sprintf(`UPDATE %s SET deleted_at = now() WHERE id = :catID`, entity.Cat{}.TableName())

    _, err := cr.db.NamedExecContext(ctx, query, map[string]interface{}{"catID": catID})
    if err != nil {
        return err
    }

    return nil
}