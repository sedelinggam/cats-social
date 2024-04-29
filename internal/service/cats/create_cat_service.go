package catsService

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/internal/entity"
	"context"
	"fmt"
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
	catData.Image = fmt.Sprintf("{%v}", strings.Join(requestData.ImageUrls, ", "))
	fmt.Println(catData.Image, "CCC")
	err = cs.catRepo.Create(ctx, catData)
	if err != nil {
		return nil, err
	}

	return &response.CreateCat{
		ID:        catData.ID,
		CreatedAt: catData.CreatedAt.Format(time.RFC3339),
	}, nil
}
