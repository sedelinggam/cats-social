package catHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"

	"github.com/gofiber/fiber/v2"
)

func (ch catHandler) GetCats(c *fiber.Ctx) error {
	var (
		req request.GetCats
		// resp *[]response.GetCats
		err error
	)
	err = c.QueryParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(response.Common{
		Message: "success",
		Data:    req,
	})
	// ctx := c.Context()
	// resp, err = ch.catService.GetCats(ctx, req)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	// return c.JSON(response.Common{
	// 	Message: "success",
	// 	Data:    resp,
	// })
}
