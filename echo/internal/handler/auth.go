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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}

	app := app.FromContext(c)

	user, err := app.Services.Auth.Register(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Register: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "register failed",
			"details": err.Error(),
		})
	}
	// return c.JSON(http.StatusCreated, user)
	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"user":   user,
	})
}

func Login(c echo.Context) error {
	var req param.Login

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}

	app := app.FromContext(c)

	token, err := app.Services.Auth.Login(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Login: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "login failed",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"token":  token,
	})
}

func Verify(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
	})
}
