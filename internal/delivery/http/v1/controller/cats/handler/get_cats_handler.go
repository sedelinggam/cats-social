package catHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (ch catHandler) GetCats(c *fiber.Ctx) error {
	var (
		req  request.GetCats
		resp *[]response.GetCats
		err  error
	)
	m := c.Queries()

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

	if m["limit"] == "" {
		req.Limit = 5
	}

	if m["offset"] == "" {
		req.Offset = 0
	}

	ctx := c.Context()
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
