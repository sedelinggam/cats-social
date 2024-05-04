package matchesService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"
	"context"
	"errors"
	"fmt"
)

func (ms matchService) ApproveMatch(ctx context.Context, requestData request.ApproveMatch) (*response.ApproveMatch, error) {
	// Get the match
	matchID := requestData.MatchID
	match, err := ms.matchRepo.GetMatchByID(ctx, matchID)

	// Check if the match is not found
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}

		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	// Get user ID
	userID := ctx.Value("user_id").(string)
	// Check if the user is not the owner of the matched cat
	if match.ReceiverID != userID {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("user is not the owner of the matched cat"))
	}

	// Check if either of the cats is already matched
	if match.UserCat.IsAlreadyMatch || match.MatchCat.IsAlreadyMatch {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("cat is already matched"))
	}

	catIds := []string{match.UserCatID, match.MatchCatID}

	// Approve the match by updating the is_already_matched field
	err = ms.catRepo.UpdateIsAlreadyMatch(ctx, catIds, true)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	fmt.Println()

	// Delete other matches
	err = ms.matchRepo.DeleteOtherMatches(ctx, match.IssuerID, match.ReceiverID, matchID)

	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.ApproveMatch{
		MatchID:    matchID,
		UserCatID:  match.UserCatID,
		MatchCatID: match.MatchCatID,
	}, nil
}
