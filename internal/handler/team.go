package handler

import (
	"net/http"

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
