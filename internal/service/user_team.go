package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/types/models/dto"
	"github.com/labstack/echo/v4"
)

type UserTeamService struct {
	repo repository.UserTeamRepository
}

func NewUserTeamService(repo repository.UserTeamRepository) *UserTeamService {
	return &UserTeamService{repo: repo}
}

func (s *UserTeamService) GetMembers(c echo.Context, teamID int) ([]dto.Member, error) {

	members, err := s.repo.GetMembers(teamID)
	if err != nil {
		c.Logger().Errorf("Service | UserTeamService | GetMembers: %v", err)
		return nil, err
	}

	dtos := make([]dto.Member, len(members))
	for i, u := range members {
		dtos[i] = dto.Member{
			UserID:   u.User.ID,
			Name:     u.User.Name,
			Email:    u.User.Email,
			TeamID:   u.Team.ID,
			TeamName: u.Team.Name,
			Role:     u.Role,
		}
	}
	return dtos, nil
}
