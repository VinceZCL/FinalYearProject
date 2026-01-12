package endpoint

import (
	"github.com/VinceZCL/FinalYearProject/internal/endpoint/middlewares"
	"github.com/VinceZCL/FinalYearProject/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e echo.Echo) {
	e.GET("/", handler.Hello)

	api := e.Group("/api")

	// api/auth/*
	auth := api.Group("/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	auth.GET("/verify", handler.Verify, middlewares.AuthMiddleware())

	// api/*
	protected := api.Group("")
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

	protected.POST("/users", handler.NewUser)

	// TODO PUT METHODS
}
