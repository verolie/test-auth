package function

import (
	"net/http"
	"test-auth/internal/utils"

	"github.com/labstack/echo/v4"
)

func ProfileHandler(c echo.Context) error {
	data := map[string]string{"message": "This is a protected route"}
	return utils.SuccessResponse(c, http.StatusOK, "Profile retrieved successfully", data)
}
