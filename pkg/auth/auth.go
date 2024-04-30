package auth

import (
	"cats-social/config"
	"cats-social/internal/entity"
	"cats-social/pkg/lumen"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// Middleware JWT function
func NewAuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.JWTSecret()),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			err = lumen.NewError(lumen.ErrUnauthorized, err)
			return lumen.FromError(err).SendResponse(c)
		},
	})
}

func GenerateToken(userData entity.User) (*string, error) {
	claims := jwt.MapClaims{
		"id":    userData.ID,
		"name":  userData.Name,
		"email": userData.Email,
		"exp":   time.Now().Add(time.Hour * 8).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JWTSecret()))
	if err != nil {
		return nil, err
	}
	return &t, nil
}
