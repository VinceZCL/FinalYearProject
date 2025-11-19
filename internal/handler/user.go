package handler

import (
	"net/http"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	app := app.FromContext(c)
	users, err := app.Services.User.GetUsers(c)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUsers: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Faied to get users"})
	}
	return c.JSON(http.StatusOK, users)
}
