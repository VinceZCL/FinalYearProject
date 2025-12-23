package endpoint

import (
	"github.com/VinceZCL/FinalYearProject/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e echo.Echo) {
	e.GET("/", handler.Hello)
	e.GET("/members/teams/:id", handler.GetMembers)
	e.GET("/users", handler.GetUsers)
	e.GET("/teams", handler.GetTeams)

	e.GET("/checkins/users/:id", handler.GetUserCheckIns)
	e.GET("/checkins/teams/:id", handler.GetTeamCheckIns)

	e.GET("/users/:id", handler.GetUser)
	e.GET("/teams/:id", handler.GetTeam)
	e.GET("/checkins/:id", handler.GetCheckIn)

	e.GET("/teams/users/:id", handler.GetUserTeams)

	e.POST("/teams", handler.NewTeam)
	e.POST("/members", handler.NewMember)
	e.POST("/checkins", handler.NewCheckIn)

	e.POST("/auth/register", handler.Register)
	e.POST("/auth/login", handler.Login)
}
