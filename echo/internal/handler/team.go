package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func GetTeams(c echo.Context) error {
	app := app.FromContext(c)

	teams, err := app.Services.Team.GetTeams(c)
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | GetTeams: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get teams"})
	}
	return c.JSON(http.StatusOK, teams)
}

func GetTeam(c echo.Context) error {
	app := app.FromContext(c)

	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid route param"})
	}

	team, err := app.Services.Team.GetTeam(c, uint(teamID))
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | GetTeam (%d): %w", teamID, err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get teams"})
	}
	return c.JSON(http.StatusOK, team)
}

func NewTeam(c echo.Context) error {
	var req param.NewTeam

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | TeamHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body"})
	}

	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | UserHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body"})
	}

	app := app.FromContext(c)

	team, err := app.Services.Team.NewTeam(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | NewTeam: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create team"})
	}

	memberReq := param.NewMember{
		UserID: team.CreatorID,
		TeamID: team.ID,
		Role:   "admin",
	}

	member, err := app.Services.UserTeam.NewMember(c, memberReq)
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | NewTeam: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create member"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"team": team, "member": member})

}
