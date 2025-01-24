package utils

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Response struct {
	App     string      `json:"app"`
	Version string      `json:"version"`
	Date    string      `json:"date"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponse(data interface{}, message string) Response {
	return Response{
		App:     "go-crud",
		Version: "1.0.0",
		Date:    time.Now().Format(time.RFC3339),
		Data:    data,
		Message: message,
	}
}

func Respond(ctx echo.Context, statusCode int, data interface{}, message string) error {
	response := NewResponse(data, message)
	return ctx.JSON(statusCode, response)
}
