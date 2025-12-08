package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/labstack/echo/v4"
)

func GetTeams(c echo.Context) error {
	app := app.FromContext(c)

	teams, err := app.Services.Team.GetTeams(c)
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | GetTeams: %w", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get teams"})
	}
	return c.JSON(http.StatusOK, teams)
}

func GetTeam(c echo.Context) error {
	app := app.FromContext(c)

	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid route param"})
	}

	team, err := app.Services.Team.GetTeam(c, teamID)
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | GetTeam (%d): %w", teamID, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get teams"})
	}
	return c.JSON(http.StatusOK, team)
}
