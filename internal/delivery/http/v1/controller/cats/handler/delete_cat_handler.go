package catHandler

import (
	"cats-social/internal/delivery/http/v1/response"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	// "fmt"
)

func (ch catHandler) DeleteCat(c *fiber.Ctx) error {
	var (
		resp *response.DeleteCat
		err  error
	)

	catID := c.Params("id")

	ctx := c.Context()

	resp, err = ch.catService.DeleteCat(ctx, catID)
	if err != nil {
		// Check if the error is due to invalid UUID input syntax
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "invalid_text_representation" {
			// Return 404 Not Found with custom message if invalid UUID
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "id tidak ditemukan",
			})
		}
	}

	return c.JSON(response.Common{
		Message: "success",
		Data:    resp,
	})
}


