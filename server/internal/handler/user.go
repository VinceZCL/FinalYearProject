package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/internal/service"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	app := app.FromContext(c)
	users, err := app.Services.User.GetUsers(c)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUsers: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "get users failed",
			"details": err.Error(),
		})
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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "invalid param",
			"details": "Invalid route param id",
		})
	}

	user, err := app.Services.User.GetUser(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetUser (%d): %w", userID, err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "get user failed",
			"details": "User not found",
		})
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
		return c.JSON(http.StatusForbidden, echo.Map{
			"status":  "failure",
			"error":   "invalid operation",
			"details": "Admin required",
		})
	}

	var req param.NewUser

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}

	app := app.FromContext(c)

	user, err := app.Services.User.NewUser(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | NewUser: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "create user failed",
			"details": err.Error(),
		})
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
		return c.JSON(http.StatusForbidden, echo.Map{
			"status":  "failure",
			"error":   "invalid operation",
			"details": "Admin required",
		})
	}

	app := app.FromContext(c)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "invalid param",
			"details": "Invalid route param id",
		})
	}

	user, err := app.Services.User.DeactivateUser(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | DeactivateUser: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "invalid param",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"user":   user,
	})
}
