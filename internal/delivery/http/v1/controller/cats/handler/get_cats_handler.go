package catHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"
	"cats-social/pkg/util"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (ch catHandler) GetCats(c *fiber.Ctx) error {
	var (
		req  request.GetCats
		resp *[]response.GetCats
		err  error
	)
	queryParams := c.Queries()

	err = c.QueryParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validator := validator.New()
	err = validator.Struct(req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	ctx := c.Context()

	shouldFilters := request.ShouldFilters{
		ID:         queryParams["id"] != "",
		Limit:      queryParams["limit"] != "",
		Offset:     queryParams["offset"] != "",
		Race:       queryParams["race"] != "",
		Sex:        queryParams["sex"] != "",
		HasMatched: queryParams["hasMatched"] != "",
		AgeInMonth: queryParams["ageInMonth"] != "",
		Owned:      queryParams["owned"] != "",
		Search:     queryParams["search"] != "",
	}

	if shouldFilters.AgeInMonth {
		num := util.ParseAgeInMonth(req.AgeInMonth)

		if num == -1 {
			return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, errors.New("invalid age in month value"))).SendResponse(c)
		}
	}

	if queryParams["limit"] == "" {
		req.Limit = 5
	}

	if queryParams["offset"] == "" {
		req.Offset = 0
	}

	ctx.SetUserValue("should_filters", shouldFilters)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	ctx.SetUserValue("user_id", userID)

	resp, err = ch.catService.GetCats(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(response.Common{
		Message: "success",
		Data:    resp,
	})
}
