package catsService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"context"
	"strings"
)

func (cs catService) GetCats(ctx context.Context, requestData request.GetCats) (*[]response.GetCats, error) {

	cats, err := cs.catRepo.GetCats(ctx, requestData)

	if err != nil {
		return nil, err
	}
	catSlice := *cats
	catsResp := []response.GetCats{}

	for i := 0; i < len(catSlice); i++ {
		cat := catSlice[i]
		catData := response.GetCats{
			ID:          cat.ID,
			Name:        cat.Name,
			Race:        cat.Race,
			Sex:         cat.Sex,
			AgeInMonth:  cat.AgeInMonth,
			Description: cat.Description,
			HasMatched:  cat.IsAlreadyMatch,
			CreatedAt:   cat.CreatedAt,
		}
		catData.ImageUrls = strings.Split(cat.Image, ",")
		catsResp = append(catsResp, catData)
	}

	return &catsResp, nil
}
