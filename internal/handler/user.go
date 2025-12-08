package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	app := app.FromContext(c)
	users, err := app.Services.User.GetUsers(c)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUsers: %w", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get users"})
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	app := app.FromContext(c)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid route param"})
	}

	user, err := app.Services.User.GetUser(c, userID)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUser (%d): %w", userID, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}
	return c.JSON(http.StatusOK, user)
}
