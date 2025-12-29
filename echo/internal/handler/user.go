package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	app := app.FromContext(c)
	users, err := app.Services.User.GetUsers(c)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUsers: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get users"})
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	app := app.FromContext(c)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid route param"})
	}

	user, err := app.Services.User.GetUser(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUser (%d): %w", userID, err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get user"})
	}
	return c.JSON(http.StatusOK, user)
}

// unused, migrate to auth.Register
func NewUser(c echo.Context) error {
	var req param.NewUser

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body"})
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body"})
	}

	app := app.FromContext(c)

	user, err := app.Services.User.NewUser(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | NewUser: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}
