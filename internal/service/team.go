package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/types/models/dto"
	"github.com/labstack/echo/v4"
)

type TeamService struct {
	repo repository.TeamRepository
}

func NewTeamService(repo repository.TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) GetTeams(c echo.Context) ([]dto.Team, error) {
	teams, err := s.repo.GetTeams()
	if err != nil {
		c.Logger().Errorf("Service | TeamService | GetTeams: %w", err)
		return nil, err
	}

	dtos := make([]dto.Team, len(teams))
	for i, u := range teams {
		dtos[i] = dto.Team{
			TeamID:      u.ID,
			TeamName:    u.Name,
			CreatorID:   u.CreatorID,
			CreatorName: u.Creator.Name,
		}
	}
	return dtos, nil
}
