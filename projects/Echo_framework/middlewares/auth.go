package middlewares

import (
	"github.com/labstack/echo/v4"
)

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username != "admin" || password != "123" {
		return false, nil
	}

	return true, nil
}
