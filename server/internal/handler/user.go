package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/internal/service"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	app := app.FromContext(c)
	users, err := app.Services.User.GetUsers(c)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUsers: %w", err)
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"users":  users,
	})
}

func GetUser(c echo.Context) error {
	app := app.FromContext(c)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Params: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	user, err := app.Services.User.GetUser(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUser (%d): %w", userID, err)
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"user":   user,
	})
}

// unused, migrate to auth.Register
func NewUser(c echo.Context) error {

	// only admin can create user
	claims := c.Get("user").(*service.Claims)
	if claims.Type != "admin" {
		return tools.ErrForbidden("Admin required")
	}

	var req param.NewUser

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	app := app.FromContext(c)

	user, err := app.Services.User.NewUser(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | NewUser: %w", err)
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"user":   user,
	})
}

func DeactivateUser(c echo.Context) error {

	// only admin can deactivate user
	claims := c.Get("user").(*service.Claims)
	if claims.Type != "admin" {
		return tools.ErrForbidden("Admin required")
	}

	app := app.FromContext(c)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Params: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	user, err := app.Services.User.DeactivateUser(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | DeactivateUser: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"user":   user,
	})
}
