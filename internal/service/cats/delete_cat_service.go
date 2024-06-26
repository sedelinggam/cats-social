package catsService

import (
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"
	"context"
	"time"
)

func (cs catService) DeleteCat(ctx context.Context, catID string) (*response.DeleteCat, error) {
	//Get user ID
	userID := ctx.Value("user_id").(string)

	// Get current time as deletion timestamp
	deletionTime := time.Now()

	// Update Cat data to mark it as deleted
	err := cs.catRepo.Delete(ctx, catID, userID)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}

		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.DeleteCat{
		ID:        catID,
		DeletedAt: deletionTime.Format(time.RFC3339),
	}, nil
}
