package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/labstack/echo/v4"
)

func GetMembers(c echo.Context) error {
	app := app.FromContext(c)

	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Params: %v", err)
		return err
	}

	members, err := app.Services.UserTeam.GetMembers(c, teamID)
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | GetMembers: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get members"})
	}
	return c.JSON(http.StatusOK, members)

}
