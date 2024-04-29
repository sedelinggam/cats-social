package usersHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"

	"github.com/gofiber/fiber/v2"
)

func (uh userHandler) Login(c *fiber.Ctx) error {
	var (
		req  request.UserLogin
		resp *response.UserAccessToken
		err  error
	)
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp, err = uh.userService.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(resp)
}
