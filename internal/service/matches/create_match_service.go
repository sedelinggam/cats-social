package matchesService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/internal/entity"
	"cats-social/pkg/lumen"
	"context"
	"time"

	"github.com/google/uuid"
)

func (ms matchService) CreateMatch(ctx context.Context, requestData request.CreateMatch) (*response.CreateMatch, error) {

	userCat, err := ms.catRepo.GetById(ctx, requestData.UserCatID)

	// Check if the user cat is not found
	if err != nil {

		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}

		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Get user ID
	userID := ctx.Value("user_id").(string)

	// Check if cat is belong to the user
	if userCat.UserID != userID {
		return nil, lumen.NewError(lumen.ErrNotFound, err)
	}

	// Check if the matched cat is not found

	matchCat, err := ms.catRepo.GetById(ctx, requestData.MatchCatID)

	if err != nil {

		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}

		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	// Check if the user cat gender is same
	if userCat.Sex == matchCat.Sex {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	// Check if the user cat is already matched
	if userCat.IsAlreadyMatch || matchCat.IsAlreadyMatch {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	// Check if same person is trying to match

	if userCat.UserID == matchCat.UserID {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	// TODO: Implement transaction

	//Create Match
	matchData := entity.Match{
		ID:         uuid.NewString(),
		MatchCatID: requestData.MatchCatID,
		UserCatID:  requestData.UserCatID,
		IssuerID:   userCat.UserID,
		ReceiverID: matchCat.UserID,
		Message:    requestData.Message,
		CreatedAt:  time.Now(),
	}

	//

	err = ms.matchRepo.Create(ctx, matchData)
	if err != nil {
		return nil, err
	}

	// TODO: Update the user cat and match cat is already matched

	return &response.CreateMatch{
		MatchCatID: matchData.MatchCatID,
		UserCatID:  matchData.UserCatID,
		Message:    matchData.Message,
	}, nil
}
