package matchesService

import (
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"
	"context"
	"strings"
	"time"
)

func (ms matchService) GetMatches(ctx context.Context, userID string) ([]response.GetMatches, error) {

	matches, err := ms.matchRepo.GetMatches(ctx, userID)

	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	if len(matches) == 0 {
		return []response.GetMatches{}, nil
	}

	var resp []response.GetMatches

	for _, match := range matches {
		match.MatchCat.Image = strings.Replace(match.MatchCat.Image, "{", "", -1)
		match.MatchCat.Image = strings.Replace(match.MatchCat.Image, "}", "", -1)
		match.UserCat.Image = strings.Replace(match.MatchCat.Image, "{", "", -1)
		match.UserCat.Image = strings.Replace(match.MatchCat.Image, "}", "", -1)
		resp = append(resp, response.GetMatches{
			ID: match.ID,
			IssuedBy: response.GetUser{
				Name:      match.Issuer.Name,
				Email:     match.Issuer.Email,
				CreatedAt: match.Issuer.CreatedAt.Format(time.RFC3339),
			},
			MatchCatDetail: response.GetCat{
				ID:          match.MatchCat.ID,
				Name:        match.MatchCat.Name,
				Race:        match.MatchCat.Race,
				Sex:         match.MatchCat.Sex,
				Description: match.MatchCat.Description,
				AgeInMonth:  match.MatchCat.AgeInMonth,
				ImageUrls:   strings.Split(match.MatchCat.Image, ","), // split image urls
				HasMatched:  match.MatchCat.IsAlreadyMatch,
				CreatedAt:   match.MatchCat.CreatedAt.Format(time.RFC3339),
			},
			UserCatDetail: response.GetCat{
				ID:          match.UserCat.ID,
				Name:        match.UserCat.Name,
				Race:        match.UserCat.Race,
				Sex:         match.UserCat.Sex,
				Description: match.UserCat.Description,
				AgeInMonth:  match.UserCat.AgeInMonth,
				ImageUrls:   strings.Split(match.UserCat.Image, ","), // split image urls
				HasMatched:  match.UserCat.IsAlreadyMatch,
				CreatedAt:   match.UserCat.CreatedAt.Format(time.RFC3339),
			},
			Message:   match.Message,
			CreatedAt: match.CreatedAt.Format(time.RFC3339),
		})
	}

	return resp, nil
}
