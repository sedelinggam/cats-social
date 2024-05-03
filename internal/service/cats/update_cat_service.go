package catsService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/internal/entity"
	"cats-social/pkg/lumen"
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func (cs catService) UpdateCat(ctx context.Context, requestData request.UpdateCat) (*response.UpdateCat, error) {
	//Password Hash
	var (
		err error
	)

	//Get user ID
	userID := ctx.Value("user_id").(string)
	//Create Cat
	catData := entity.Cat{
		ID:          requestData.ID,
		UserID:      userID,
		Name:        requestData.Name,
		Race:        requestData.Race,
		Sex:         requestData.Sex,
		AgeInMonth:  requestData.AgeInMonth,
		Description: requestData.Description,
		UpdatedAt:   time.Now(),
	}

	//Check if url
	for _, v := range requestData.ImageUrls {
		_, err := url.ParseRequestURI(v)
		if err != nil {
			lumen.NewError(lumen.ErrBadRequest, err)
		}
		catData.Image = fmt.Sprintf("{%v}", strings.Join(requestData.ImageUrls, ", "))
	}

	err = cs.catRepo.Update(ctx, catData)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.UpdateCat{
		ID:        catData.ID,
		UpdatedAt: catData.UpdatedAt.Format(time.RFC3339),
	}, nil
}
