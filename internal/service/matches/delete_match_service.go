package matchesService

import (
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"
	"context"
	"time"
)

func (cs matchService) DeleteMatch(ctx context.Context, id string) (*response.DeleteCat, error) {

	//Get user ID
	userID := ctx.Value("user_id").(string)

	//Check if cat is match
	cat, err := cs.matchRepo.GetMatches(ctx, id)
	if cat != nil {
		lumen.NewError(lumen.ErrBadRequest, err)
	}

	// Get current time as deletion timestamp
	deletionTime := time.Now()

	// Update Cat data to mark it as deleted
	err = cs.matchRepo.Delete(ctx, id, userID)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		lumen.NewError(lumen.ErrBadRequest, err)
	}

	return &response.DeleteCat{
		ID:        id,
		DeletedAt: deletionTime.Format(time.RFC3339),
	}, nil
}
