package handler

import (
	"net/http"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

func NewComment(c echo.Context) error {
	var req param.NewComment
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Handler | CommentHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}
	if err := req.Validate(); err != nil {
		c.Logger().Errorf("Handler | CommentHandler | Invalid Request: %w", err)
		return tools.ErrBadRequest(err.Error())
	}

	app := app.FromContext(c)
	err := app.Services.Comment.NewComment(c, req)
	if err != nil {
		c.Logger().Errorf("Handler | CheckInHandler | NewCheckIn: %w", err)
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"status": "success",
	})
}
