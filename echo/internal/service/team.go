package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/types/model"
	"github.com/VinceZCL/FinalYearProject/types/model/dto"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
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
			ID:          u.ID,
			Name:        u.Name,
			CreatorID:   u.CreatorID,
			CreatorName: u.Creator.Name,
		}
	}
	return dtos, nil
}

func (s *TeamService) GetTeam(c echo.Context, teamID uint) (*dto.Team, error) {
	team, err := s.repo.GetTeam(teamID)
	if err != nil {
		c.Logger().Errorf("Service | TeamService | GetTeam (%d): %w", teamID, err)
		return nil, err
	}

	dto := &dto.Team{
		ID:          team.ID,
		Name:        team.Name,
		CreatorID:   team.CreatorID,
		CreatorName: team.Creator.Name,
	}
	return dto, nil
}

func (s *TeamService) NewTeam(c echo.Context, req param.NewTeam) (*dto.Team, error) {
	input := model.Team{
		Name:      req.Name,
		CreatorID: req.CreatorID,
	}

	team, err := s.repo.NewTeam(input)
	if err != nil {
		c.Logger().Errorf("Service | TeamService | NewTeam: %w", err)
		return nil, err
	}

	dto := &dto.Team{
		ID:          team.ID,
		Name:        team.Name,
		CreatorID:   team.CreatorID,
		CreatorName: team.Creator.Name,
	}
	return dto, nil
}
