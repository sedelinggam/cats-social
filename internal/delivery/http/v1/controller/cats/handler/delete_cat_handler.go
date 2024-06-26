package catHandler

import (
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	// "fmt"
)

func (ch catHandler) DeleteCat(c *fiber.Ctx) error {
	var (
		resp *response.DeleteCat
		err  error
	)

	catID := c.Params("id")

	ctx := c.Context()

	//Get jwt user ID
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	ctx.SetUserValue("user_id", userID)

	resp, err = ch.catService.DeleteCat(ctx, catID)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(response.Common{
		Message: "success",
		Data:    resp,
	})
}
