package handler

import (
	"net/http"
	"strconv"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/labstack/echo/v4"
)

func GetUserCheckIns(c echo.Context) error {
	app := app.FromContext(c)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Params: %w", err)
		return err
	}

	checkIns, err := app.Services.CheckIn.GetUserCheckIns(c, userID)
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | GetUserCheckIns: %w", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get User Check Ins"})
	}

	return c.JSON(http.StatusOK, checkIns)
}

func GetTeamCheckIns(c echo.Context) error {
	app := app.FromContext(c)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Params: %w", err)
		return err
	}

	checkIns, err := app.Services.CheckIn.GetUserCheckIns(c, userID)
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | GetTeamCheckIns: %w", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get Team Check Ins"})
	}

	return c.JSON(http.StatusOK, checkIns)
}
