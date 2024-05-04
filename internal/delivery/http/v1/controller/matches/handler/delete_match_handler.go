package matchHandler

import (
	"cats-social/internal/delivery/http/v1/response"
	"cats-social/pkg/lumen"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	// "fmt"
)

func (mh matchHandler) DeleteCat(c *fiber.Ctx) error {
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

	resp, err = mh.matchService.DeleteMatch(ctx, catID)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(response.Common{
		Message: "success",
		Data:    resp,
	})
}
