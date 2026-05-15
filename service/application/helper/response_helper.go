package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"stock-watchlist/application/errs"
	"stock-watchlist/domain/model/dto"
	"stock-watchlist/domain/model/message"

	"github.com/gofiber/fiber/v2"
)


func ClientSuccessJSONResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(dto.SuccessResponse{
		Message: "success",
		Data: data,
	})
}

func ClientErrorJSONResponse(ctx *fiber.Ctx, err error) error {
	var appErr *errs.AppError
	var clientErr *dto.ErrorResponse
	var unmarsharlTypeError *json.UnmarshalTypeError

	if errors.As(err, &appErr) {
		clientErr = &dto.ErrorResponse{
			Message: appErr.Message,
			Data: appErr.Data,
		}
		return ctx.Status(appErr.StatusCode).JSON(clientErr)
	}

	if errors.As(err, &unmarsharlTypeError) {
		clientErr = &dto.ErrorResponse{
			Message: "invalid type for field " + unmarsharlTypeError.Field + ": expected " + unmarsharlTypeError.Type.String(),
		}
		return ctx.Status(http.StatusBadRequest).JSON(clientErr)
	}
	
	clientErr = &dto.ErrorResponse{
		Message: message.ErrInternalServer.Message,
	}

	return ctx.Status(http.StatusInternalServerError).JSON(clientErr)
}