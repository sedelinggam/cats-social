package catHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (ch catHandler) UpdateCat(c *fiber.Ctx) error {
	var (
		req  request.UpdateCat
		resp *response.UpdateCat
		err  error
	)
	err = c.BodyParser(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	// Create a new validator instance
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	//Get jwt user ID
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	ctx := c.Context()
	ctx.SetUserValue("user_id", userID)

	if c.Params("id") != "" {
		req.ID = c.Params("id")
	}
	resp, err = ch.catService.UpdateCat(ctx, req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.Status(fiber.StatusOK).JSON(response.Common{
		Message: "success",
		Data:    resp,
	})
}
