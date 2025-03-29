package function

import (
	"net/http"
	"test-auth/internal/utils"
	"test-auth/internal/utils/middleware"
	"test-auth/model"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	if err := c.Validate(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Username and password are required")
	}

	token, err := middleware.GenerateJWT(req.Username)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
	}

	return utils.SuccessResponse(c, http.StatusOK, "Login successful", map[string]string{"token": token})
}
