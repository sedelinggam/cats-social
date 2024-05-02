package catHandler

import (
	"cats-social/internal/delivery/http/v1/response"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func (ch catHandler) DeleteCat(c *fiber.Ctx) error {
	var (
		resp *response.DeleteCat
		err  error
	)

	catID := c.Params("id")

	ctx := c.Context()

	// Delete cat using service
	resp, err = ch.catService.DeleteCat(ctx, catID)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrBadRequest
	}

	return c.JSON(resp)
}