package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"golang_projects/projects/Echo_framework/models"
)

func Hello(c echo.Context) error {
	x := &models.X{
		Text: "Hello World!",
	}
	return c.JSON(http.StatusOK, x)
}