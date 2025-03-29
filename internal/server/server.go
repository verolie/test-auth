package server

import (
	"test-auth/internal/function"
	"test-auth/internal/utils/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitServer() *echo.Echo {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	handlerServer(e)

	return e
}

func handlerServer(e *echo.Echo) {
	e.POST("/login", function.LoginHandler)

	e.GET("/profile", function.ProfileHandler, middleware.AuthMiddleware)
}
