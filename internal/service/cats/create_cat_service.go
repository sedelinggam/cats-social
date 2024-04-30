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

	"github.com/google/uuid"
)

func (cs catService) CreateCat(ctx context.Context, requestData request.CreateCat) (*response.CreateCat, error) {
	//Password Hash
	var (
		err error
	)

	//Get user ID
	userID := ctx.Value("user_id").(string)
	//Create Cat
	catData := entity.Cat{
		ID:             uuid.New().String(),
		UserID:         userID,
		Name:           requestData.Name,
		Race:           requestData.Race,
		Sex:            requestData.Sex,
		AgeInMonth:     requestData.AgeInMonth,
		Description:    requestData.Description,
		IsAlreadyMatch: false,
		CreatedAt:      time.Now(),
	}
	//Check if url
	for _, v := range requestData.ImageUrls {
		_, err := url.ParseRequestURI(v)
		if err != nil {
			lumen.NewError(lumen.ErrBadRequest, err)
		}
	}
	catData.Image = fmt.Sprintf("{%v}", strings.Join(requestData.ImageUrls, ", "))
	err = cs.catRepo.Create(ctx, catData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.CreateCat{
		ID:        catData.ID,
		CreatedAt: catData.CreatedAt.Format(time.RFC3339),
	}, nil
}
