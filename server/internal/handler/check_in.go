package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func GetUserCheckIns(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "invalid param",
			"details": "Invalid route param id",
		})
	}
	date := ""
	dateParam := c.QueryParam("date")
	if dateParam != "" {
		datetime, err := time.Parse(time.DateOnly, dateParam)
		date = datetime.Format(time.DateOnly)
		if err != nil {
			c.Logger().Errorf("Handler | CheckInHandler | Invalid Params: %w", err)
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "failure",
				"error":   "invalid param",
				"details": "Invalid route param date",
			})
		}
	}

	app := app.FromContext(c)
	checkIns, err := app.Services.CheckIn.GetUserCheckIns(c, uint(userID), date)
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | GetUserCheckIns: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "get user check ins failed",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":   "success",
		"checkIns": checkIns,
	})
}

func GetTeamCheckIns(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "invalid param",
			"details": "Invalid route param id",
		})
	}
	date := ""
	dateParam := c.QueryParam("date")
	if dateParam != "" {
		datetime, err := time.Parse(time.DateOnly, dateParam)
		date = datetime.Format(time.DateOnly)
		if err != nil {
			c.Logger().Errorf("Handler | CheckInHandler | Invalid Params: %w", err)
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "failure",
				"error":   "invalid param",
				"details": "Invalid route param date",
			})
		}
	}

	app := app.FromContext(c)
	checkIns, err := app.Services.CheckIn.GetTeamCheckIns(c, uint(userID), date)
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | GetTeamCheckIns: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "get team check ins failed",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":   "success",
		"checkIns": checkIns,
	})
}

func GetCheckIn(c echo.Context) error {
	checkInID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Params: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "invalid param",
			"details": "Invalid route param id",
		})
	}

	app := app.FromContext(c)
	checkIn, err := app.Services.CheckIn.GetCheckIn(c, uint(checkInID))
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | GetCheckIn: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "get check in failed",
			"details": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"checkIn": checkIn,
	})
}

func NewCheckIn(c echo.Context) error {
	var req param.NewCheckIn
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}
	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | UserTeamHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}

	app := app.FromContext(c)
	checkIn, err := app.Services.CheckIn.NewCheckIn(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | NewCheckIn: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "create check in failed",
			"details": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"status":  "success",
		"checkIn": checkIn,
	})
}

func BulkCheckIn(c echo.Context) error {
	// Bind the request to BulkCheckIn
	var req param.BulkCheckIn
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "malformed JSON",
			"details": err.Error(),
		})
	}

	// Validate the bulk check-in request
	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | Invalid Request: %w", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "failure",
			"error":   "invalid check in",
			"details": err.Error(),
		})
	}

	// Get the service instance from the context
	app := app.FromContext(c)

	// Call the BulkCheckIn service
	checkIns, err := app.Services.CheckIn.BulkCheckIn(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | BulkCheckIn: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "failure",
			"error":   "bulk check-ins failed",
			"details": err.Error(),
		})
	}

	// Return the successful response with the check-ins grouped by user
	return c.JSON(http.StatusCreated, echo.Map{
		"status":   "success",
		"checkIns": checkIns,
	})
}
