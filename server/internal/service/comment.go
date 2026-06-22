package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) NewComment(c echo.Context, req param.NewComment) error {
	input := model.Comment{
		UserID:    req.UserID,
		CheckinID: req.CheckinID,
		TeamID:    req.TeamID,
		Item:      req.Item,
	}

	err := s.repo.NewComment(input)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | NewCheckIn: %w", err)
		return tools.ErrInternal("database failure", err.Error())
	}

	return nil
}
