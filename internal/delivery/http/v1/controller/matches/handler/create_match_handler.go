package matchHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"

	"github.com/gofiber/fiber/v2"
)

func (mh matchHandler) CreateMatch(c *fiber.Ctx) error {
	var (
		req  request.CreateMatch
		resp *response.CreateMatch
		err  error
	)
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	ctx := c.Context()
	resp, err = mh.matchService.CreateMatch(ctx, req)
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
