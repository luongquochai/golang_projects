package handler
import (
	"net/http"
	"github.com/labstack/echo/v4"
	"golang_projects/projects/Echo_framework/models"
)

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: "77777",
	})
}