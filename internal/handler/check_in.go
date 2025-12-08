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

	checkIns, err := app.Services.CheckIn.GetUserCheckIns(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | GetUserCheckIns: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get User Check Ins"})
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

	checkIns, err := app.Services.CheckIn.GetUserCheckIns(c, uint(userID))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | GetTeamCheckIns: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get Team Check Ins"})
	}

	return c.JSON(http.StatusOK, checkIns)
}

func GetCheckIn(c echo.Context) error {
	app := app.FromContext(c)

	checkInID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid route param"})
	}

	checkIn, err := app.Services.CheckIn.GetCheckIn(c, uint(checkInID))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | GetCheckIn: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get Check In"})
	}
	return c.JSON(http.StatusOK, checkIn)
}
