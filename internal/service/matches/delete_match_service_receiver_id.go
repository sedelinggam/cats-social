package matchesService

import (
	"cats-social/pkg/lumen"
	"context"
)

func (cs matchService) DeleteMatchReceiverID(ctx context.Context, id string) error {

	//Get user ID
	userID := ctx.Value("user_id").(string)

	//Check if cat is match
	cat, err := cs.matchRepo.GetMatches(ctx, id)
	if cat != nil {
		lumen.NewError(lumen.ErrBadRequest, err)
	}

	// Update Cat data to mark it as deleted
	err = cs.matchRepo.DeleteReceiverID(ctx, id, userID)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return lumen.NewError(lumen.ErrNotFound, err)
		}
		lumen.NewError(lumen.ErrBadRequest, err)
	}

	return nil
}
