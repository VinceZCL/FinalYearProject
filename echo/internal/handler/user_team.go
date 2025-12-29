package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func GetMembers(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid route param"})
	}

	app := app.FromContext(c)

	members, err := app.Services.UserTeam.GetMembers(c, uint(teamID))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | GetMembers: %w", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get members"})
	}
	return c.JSON(http.StatusOK, members)

}

func NewMember(c echo.Context) error {
	var req param.NewMember
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body"})
	}
	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON Body"})
	}

	app := app.FromContext(c)

	member, err := app.Services.UserTeam.NewMember(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | NewTeam: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create member"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"member": member})
}

func GetUserTeams(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid route param"})
	}

	app := app.FromContext(c)

	members, err := app.Services.UserTeam.GetUserTeams(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | GetUserTeams: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get user's teams"})
	}
	return c.JSON(http.StatusOK, echo.Map{"member": members})
}
