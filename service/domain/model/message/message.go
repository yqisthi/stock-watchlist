package message

import (
	"net/http"
	"stock-watchlist/application/errs"
)

type Message struct {
	Code int
	Message string
}

func (m *Message) GetError() error {
	return &errs.AppError{
		StatusCode: m.Code,
		Message: m.Message,
	}
}

func newMessage(code int, message string) *Message {
	return &Message{
		Code: code,
		Message: message,
	}
}

var (
	// 400 Bad Request
	ErrParseRequest = newMessage(http.StatusBadRequest, "failed to parse request body")
	
	// 401 Unauthorized
	ErrInvalidCredentials = newMessage(http.StatusUnauthorized, "invalid credentials")

	// 404 Not Found
	ErrUserNotFound       = newMessage(http.StatusNotFound, "user not found")
	ErrRecordNotFound	 = newMessage(http.StatusNotFound, "record not found")

	// 500 Internal Server Error
	ErrInternalServer     = newMessage(http.StatusInternalServerError, "internal server error")
)
