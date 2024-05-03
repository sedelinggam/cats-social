package catHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"
	"cats-social/pkg/util"
	"strconv"

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
	validator := validator.New()
	ctx := c.Context()

	req = request.GetCats{
		Search: queryParams["search"],
	}

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

	if shouldFilters.ID {
		err := validator.Var(queryParams["id"], "uuid")
		if err != nil {
			shouldFilters.ID = false
		} else {
			req.ID = queryParams["id"]
		}
	}

	if shouldFilters.Race {
		err := validator.Var(queryParams["race"], "oneof=Persian 'Maine Coon' Ragdoll Bengal Sphynx 'British Shorthair' Abyssinian 'Scottish Fold' Birman")
		if err != nil {
			shouldFilters.Race = false
		} else {
			req.Race = queryParams["race"]
		}
	}

	if shouldFilters.Sex {
		err := validator.Var(queryParams["sex"], "oneof=male female")
		if err != nil {
			shouldFilters.Sex = false
		} else {
			req.Sex = queryParams["sex"]
		}
	}

	if shouldFilters.HasMatched {
		err := validator.Var(queryParams["hasMatched"], "boolean")
		if err != nil {
			shouldFilters.HasMatched = false
		} else {
			val, _ := strconv.ParseBool(queryParams["hasMatched"])
			req.HasMatched = val
		}
	}

	if shouldFilters.Owned {
		err := validator.Var(queryParams["owned"], "boolean")
		if err != nil {
			shouldFilters.Owned = false
		} else {
			val, _ := strconv.ParseBool(queryParams["owned"])
			req.Owned = val
		}
	}

	if shouldFilters.AgeInMonth {
		num := util.ParseAgeInMonth(queryParams["ageInMonth"])

		if num == -1 {
			shouldFilters.AgeInMonth = false
		} else {
			req.AgeInMonth = queryParams["ageInMonth"]
		}
	}

	if shouldFilters.Limit {
		val, _ := strconv.ParseInt(queryParams["limit"], 10, 32)
		req.Limit = int32(val)
	} else {
		req.Limit = 5
	}

	if shouldFilters.Offset {
		val, _ := strconv.ParseInt(queryParams["offset"], 10, 32)
		req.Offset = int32(val)
	} else {
		req.Offset = 0
	}

	ctx.SetUserValue("should_filters", shouldFilters)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	ctx.SetUserValue("user_id", userID)

	resp, err = ch.catService.GetCats(ctx, req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(response.Common{
		Message: "success",
		Data:    resp,
	})
}
