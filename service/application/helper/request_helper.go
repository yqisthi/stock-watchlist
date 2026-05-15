package helper

import (
	"stock-watchlist/application/errs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseRequest[T interface{}](ctx *fiber.Ctx) (*T, error) {
	var request T

	err := ctx.BodyParser(&request);
	return &request, err
}

func ValidateRequest(data interface{}, validator *validator.Validate) error {
	err := validator.Struct(data)

	if err == nil {
		return err
	}

	return errs.NewValidationError(err)
}