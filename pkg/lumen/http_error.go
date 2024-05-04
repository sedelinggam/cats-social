package lumen

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	Status  int
	Message string
}

func FromError(err error) APIError {
	var apiError APIError
	var svcError Error
	if errors.As(err, &svcError) {
		if svcError.appError == nil {
			apiError.Message = "generic error"
		}
		apiError.Message = svcError.appError.Error()
		svcErr := svcError.SvcError()
		switch svcErr {
		case ErrBadRequest:
			apiError.Status = fiber.StatusBadRequest
		case ErrUnauthorized:
			apiError.Status = fiber.StatusUnauthorized
		case ErrNotFound:
			apiError.Status = fiber.StatusNotFound
		case ErrConflict:
			apiError.Status = fiber.StatusConflict
		case ErrInternalFailure:
			apiError.Status = fiber.StatusInternalServerError
		}

	}
	return apiError
}

func (apiErr APIError) SendResponse(c *fiber.Ctx) error {
	return c.Status(apiErr.Status).JSON(fiber.Map{
		"error": apiErr.Message,
	})
}
