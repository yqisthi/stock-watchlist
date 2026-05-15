package errs

import (
	"errors"
	"net/http"
	"stock-watchlist/domain/model/dto"
	"strings"

	"github.com/ettle/strcase"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type AppError struct {
	StatusCode int
	Message string
	Data any
	OriginalError error
}

func (e *AppError) Error() string {
	return e.Message
}

func NewDatabaseError(err error) *AppError {
	var pgError *pgconn.PgError

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &AppError{
			StatusCode: http.StatusNotFound,
			Message: "record not found",
			OriginalError: err,
		}
	} else if errors.As(err, &pgError) {
		switch pgError.Code {
		case "23505": // unique_violation
		  fieldName := extractFieldNameFromConstraint(pgError.ConstraintName)
			return &AppError{
				StatusCode: http.StatusBadRequest,
				Message: "duplicate record",
				Data:  fieldName + " already exists",
				OriginalError: err,
			}
		default:
			return &AppError{
				StatusCode: http.StatusInternalServerError,
				Message: "database error",
				OriginalError: err,
			}
		}
	}

	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message: "internal server error",
		OriginalError: err,
	}
}

func NewValidationError(err error) *AppError {
	var validationError validator.ValidationErrors
	appError := &AppError{
		StatusCode: http.StatusBadRequest,
		Message: "validation error",
		OriginalError: err,
	}

	if errors.As(err, &validationError) {
		var errs []dto.FieldError
		for _, fieldErr := range validationError {
			var errorMsg string
			switch fieldErr.Tag() {
			case "required":
				errorMsg = "field is required"
			case "email":
				errorMsg = "field must be a valid email"
			case "min":
				errorMsg = "field must be at least " + fieldErr.Param() + " characters long"
			case "max":
				errorMsg = "field must be at most " + fieldErr.Param() + " characters long"
			default:
				errorMsg = fieldErr.Error()
			}

			errs = append(errs, dto.FieldError{
				Field: strcase.ToSnake(fieldErr.Field()),
				Error: errorMsg,
			})
		}

		appError.Data = errs
	}

	return appError
}

func extractFieldNameFromConstraint(constraintName string) string {
	// Ex: "unique_username" -> "username"
	parts := strings.Split(constraintName, "_")
	if len(parts) > 1 {
		return strings.Join(parts[1:], "_")
	}
	return constraintName
}