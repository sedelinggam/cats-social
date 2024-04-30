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
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}
	// Create a new validator instance
	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	resp, err = uh.userService.Register(c.Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}
	return c.Status(fiber.StatusCreated).JSON(response.Common{
		Message: "User registered successfully",
		Data:    resp,
	})
}
