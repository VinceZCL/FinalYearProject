package endpoint

import (
	"github.com/VinceZCL/FinalYearProject/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e echo.Echo) {
	e.GET("/", handler.Hello)
	e.GET("/members/:id", handler.GetMembers)
	e.GET("/users", handler.GetUsers)
	e.GET("/teams", handler.GetTeams)
}
