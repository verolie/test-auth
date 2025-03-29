package utils

import (
	"test-auth/model"

	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	response := model.SuccessResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	return c.JSON(statusCode, response)
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	response := model.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}

	return c.JSON(statusCode, response)
}
