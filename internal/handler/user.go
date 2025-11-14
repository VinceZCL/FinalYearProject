package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"scrum.com/app"
)

func GetMembers(c echo.Context) error {
	app := app.FromContext(c)

	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Params: %v", err)
		return err
	}

	members, err := app.Services.User.GetMembers(c, teamID)
	if err != nil {
		c.Logger().Errorf("Handler | UserHandler | GetMembers: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get members"})
	}
	return c.JSON(http.StatusOK, members)

}
