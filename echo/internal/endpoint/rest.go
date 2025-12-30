package endpoint

import (
	"github.com/VinceZCL/FinalYearProject/internal/endpoint/middlewares"
	"github.com/VinceZCL/FinalYearProject/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e echo.Echo) {
	e.GET("/", handler.Hello)
	e.POST("/auth/register", handler.Register)
	e.POST("/auth/login", handler.Login)

	protected := e.Group("")
	protected.Use(middlewares.AuthMiddleware())

	// GET ALL
	protected.GET("/users", handler.GetUsers)
	protected.GET("/teams", handler.GetTeams)

	// GET SINGLE
	protected.GET("/users/:id", handler.GetUser)
	protected.GET("/teams/:id", handler.GetTeam)
	protected.GET("/checkins/:id", handler.GetCheckIn)

	// GET TEAM RELATED
	protected.GET("/members/teams/:id", handler.GetMembers)
	protected.GET("/teams/users/:id", handler.GetUserTeams)

	// GET CHECKINS
	protected.GET("/checkins/users/:id", handler.GetUserCheckIns)
	protected.GET("/checkins/teams/:id", handler.GetTeamCheckIns)

	// CREATE
	protected.POST("/teams", handler.NewTeam)
	protected.POST("/members", handler.NewMember)
	protected.POST("/checkins", handler.NewCheckIn)
	protected.POST("/checkins/bulk", handler.BulkCheckIn)

	// TODO PUT METHODS
}
