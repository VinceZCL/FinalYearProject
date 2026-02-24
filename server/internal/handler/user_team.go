package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func GetMembers(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Params: %w", err)
		return tools.ErrBadRequest("Invalid route param")
	}

	app := app.FromContext(c)

	members, err := app.Services.UserTeam.GetMembers(c, uint(teamID))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | GetMembers: %w", err)
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"members": members,
	})

}

func NewMember(c echo.Context) error {
	var req param.NewMember
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}
	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	app := app.FromContext(c)

	member, err := app.Services.UserTeam.NewMember(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | NewTeam: %w", err)
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
		"member": member,
	})
}

func GetUserTeams(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Params: %w", err)
		return tools.ErrBadRequest("Invalid route param")
	}

	app := app.FromContext(c)

	members, err := app.Services.UserTeam.GetUserTeams(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | GetUserTeams: %w", err)
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"members": members,
	})
}
