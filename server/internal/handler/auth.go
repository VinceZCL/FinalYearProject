package handler

import (
	"net/http"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var req param.NewUser

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	app := app.FromContext(c)

	user, err := app.Services.Auth.Register(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Register: %w", err)
		return err
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
		return tools.ErrBadRequest(err.Error())
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	app := app.FromContext(c)

	token, err := app.Services.Auth.Login(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | AuthHandler | Login: %w", err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"token":  token,
	})
}

func Verify(c echo.Context) error {

	claims := c.Get("user")

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"claims": claims,
	})
}
