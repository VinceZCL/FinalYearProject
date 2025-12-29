package handler

import (
	"net/http"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var req param.NewUser

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body", "details": err.Error()})
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body", "details": err.Error()})
	}

	app := app.FromContext(c)

	user, err := app.Services.Auth.Register(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Register: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user", "details": err.Error()})
	}
	// return c.JSON(http.StatusCreated, user)
	return c.JSON(http.StatusCreated, echo.Map{"Status": "Success", "User": user})
}

func Login(c echo.Context) error {
	var req param.Login

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body", "details": err.Error()})
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body", "details": err.Error()})
	}

	app := app.FromContext(c)

	token, err := app.Services.Auth.Login(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Login: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"Status": "Failure", "error": "Failed to login", "details": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"Status": "Success", "Token": token})
}
