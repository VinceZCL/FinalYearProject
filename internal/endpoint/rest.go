package endpoint

import (
	"github.com/labstack/echo/v4"
	"scrum.com/internal/handler"
)

func RegisterRoutes(e echo.Echo) {
	e.GET("/", handler.Hello)
	e.GET("/members/:id", handler.GetMembers)
}
