package matchHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (mh matchHandler) ApproveMatch(c *fiber.Ctx) error {
	var (
		req  request.ApproveMatch
		resp *response.ApproveMatch
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

	resp, err = mh.matchService.ApproveMatch(ctx, req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.Common{
		Message: "success",
		Data:    resp,
	})
}
