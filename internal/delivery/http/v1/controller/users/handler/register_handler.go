package usersHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"

	"github.com/gofiber/fiber/v2"
)

func (uh userHandler) Register(c *fiber.Ctx) error {
	var (
		req  request.UserRegister
		resp *response.UserAccessToken
		err  error
	)
	err = c.BodyParser(req)
	if err != nil {
		return fiber.ErrBadGateway
	}
	resp, err = uh.userService.Register(c.Context(), req)
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.JSON(resp)
}
