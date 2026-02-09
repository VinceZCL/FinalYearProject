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
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "get teams failed",
			"details": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"teams":  teams,
	})
}

func GetTeam(c echo.Context) error {
	app := app.FromContext(c)

	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "invalid param",
			"details": "Invalid route param id",
		})
	}

	team, err := app.Services.Team.GetTeam(c, uint(teamID))
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | GetTeam (%d): %w", teamID, err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "get teams failed",
			"details": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"team":   team,
	})
}

func NewTeam(c echo.Context) error {
	var req param.NewTeam

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | TeamHandler | Invalid Request: %w", err)
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

	team, err := app.Services.Team.NewTeam(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | NewTeam: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "create team failed",
			"details": "Team already exist",
		})
	}

	memberReq := param.NewMember{
		UserID: team.CreatorID,
		TeamID: team.ID,
		Role:   "admin",
	}

	member, err := app.Services.UserTeam.NewMember(c, memberReq)
	if err != nil {
		c.Logger().Errorf("Handler | TeamHandler | NewTeam: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "create member failed",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"team":   team,
		"member": member,
	})

}
