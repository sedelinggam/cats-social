package usersHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (uh userHandler) Register(c *fiber.Ctx) error {
	var (
		req  request.UserRegister
		resp *response.UserAccessToken
		err  error
	)
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Create a new validator instance
	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err = uh.userService.Register(c.Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}
	return c.JSON(resp)
}
