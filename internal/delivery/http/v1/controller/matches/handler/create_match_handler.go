package matchHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

	//Get jwt user ID
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	ctx := c.Context()
	ctx.SetUserValue("user_id", userID)
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
