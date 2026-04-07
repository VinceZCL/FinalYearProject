package endpoint

import (
	"os"
	"strings"

	"github.com/VinceZCL/FinalYearProject/app/config"
	"github.com/labstack/echo/v4"
)

func setupStatic(e echo.Echo) {
	clientDir := config.Get().Client.Dir
	if clientDir == "" {
		clientDir = "/app/client/browser"
	}

	e.Static("/static", clientDir+"/static")
	e.Static("/*", clientDir)

	e.GET("/", func(c echo.Context) error {
		return c.File(clientDir + "/index.html")
	})

	e.GET("/*", func(c echo.Context) error {
		path := c.Param("*")

		if len(path) >= 7 && path[:7] == "static/" {
			return echo.ErrNotFound
		}

		if len(path) >= 4 && strings.ToLower(path[:4]) == "api/" {
			return echo.ErrNotFound
		}

		if containsDot(path) {
			filePath := clientDir + "/" + path
			if fileExists(filePath) {
				return c.File(filePath)
			}
			return echo.ErrNotFound
		}

		return c.File(clientDir + "/index.html")
	})
}

func containsDot(path string) bool {
	for i := 0; i < len(path); i++ {
		if path[i] == '.' {
			return true
		}
	}
	return false
}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
