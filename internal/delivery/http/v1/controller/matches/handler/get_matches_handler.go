package matchHandler

import (
	"cats-social/internal/delivery/http/v1/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (mh matchHandler) GetMatches(c *fiber.Ctx) error {
	var (
		resp []response.GetMatches
		err  error
	)

	//Get jwt user ID
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	ctx := c.Context()

	resp, err = mh.matchService.GetMatches(ctx, userID)
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
