package catsRepository

import (
	"context"
	"fmt"
	"strings"
)

func (cr catRepository) UpdateIsAlreadyMatch(ctx context.Context, ids []string, isAlreadyMatch bool) error {

	if len(ids) == 0 {
		return nil // No IDs provided, nothing to update
	}

	// Prepare the SQL query with placeholders for the IDs
	query := `UPDATE cats SET is_already_matched = $1 WHERE id IN (`
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
	}
	query += strings.Join(placeholders, ",") + ")"

	// Prepare the query arguments
	args := make([]interface{}, len(ids)+1)
	args[0] = isAlreadyMatch
	for i, id := range ids {
		args[i+1] = id
	}

	tx := cr.db.MustBegin()

	// Execute the update query
	_, err := cr.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
