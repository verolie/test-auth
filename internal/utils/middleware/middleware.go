package middleware

import (
	"net/http"
	"test-auth/internal/utils"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")
		if tokenStr == "" {
			return utils.ErrorResponse(c, http.StatusUnauthorized, "Missing token")
		}

		token, err := VerifyJWT(tokenStr)
		if err != nil || !token.Valid {
			return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		}

		return next(c)
	}
}
