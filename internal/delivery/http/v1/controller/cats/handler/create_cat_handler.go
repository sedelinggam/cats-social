package catHandler

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (ch catHandler) CreateCat(c *fiber.Ctx) error {
	var (
		req  request.CreateCat
		resp *response.CreateCat
		err  error
	)
	err = c.BodyParser(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	ctx := c.Context()
	ctx.SetUserValue("user_id", userID)
	resp, err = ch.catService.CreateCat(ctx, req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}
	return c.JSON(response.Common{
		Message: "success",
		Data:    resp,
	})
}
