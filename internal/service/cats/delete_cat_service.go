package catsService

import (
	"cats-social/internal/delivery/http/v1/response"
	"context"
	"time"
)

func (cs catService) DeleteCat(ctx context.Context, catID string) (*response.DeleteCat, error) {
	// Get current time as deletion timestamp
	deletionTime := time.Now()

	// Update Cat data to mark it as deleted
	err := cs.catRepo.Delete(ctx, catID)
	if err != nil {
		return nil, err
	}

	return &response.DeleteCat{
		ID:        catID,
		DeletedAt: deletionTime.Format(time.RFC3339),
	}, nil
}
