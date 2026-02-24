package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model"
	"github.com/VinceZCL/FinalYearProject/types/model/dto"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

type UserTeamService struct {
	repo repository.UserTeamRepository
}

func NewUserTeamService(repo repository.UserTeamRepository) *UserTeamService {
	return &UserTeamService{repo: repo}
}

func (s *UserTeamService) GetMembers(c echo.Context, teamID uint) ([]dto.Member, error) {

	members, err := s.repo.GetMembers(teamID)
	if err != nil {
		c.Logger().Errorf("Service | UserTeamService | GetMembers: %w", err)
		return nil, tools.ErrInternal("database failure", err.Error())
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

func (s *UserTeamService) NewMember(c echo.Context, req param.NewMember) (*dto.Member, error) {

	claims := c.Get("user").(Claims)

	admin, err := s.repo.IsTeamAdmin(claims.UserID, req.TeamID)
	if err != nil {
		return nil, tools.ErrInternal("database failure", err.Error())
	}
	if !admin {
		return nil, tools.ErrForbidden("Team Admin required")
	}

	input := model.UserTeam{
		UserID: req.UserID,
		TeamID: req.TeamID,
		Role:   req.Role,
	}

	member, err := s.repo.NewMember(input)
	if err != nil {
		c.Logger().Errorf("Service | UserTeamService | NewMember: %w", err)
		return nil, tools.ErrInternal("database failure", err.Error())
	}
	dto := &dto.Member{
		UserID:   member.UserID,
		Name:     member.User.Name,
		Email:    member.User.Email,
		TeamID:   member.TeamID,
		TeamName: member.Team.Name,
		Role:     member.Role,
	}
	return dto, nil
}

func (s *UserTeamService) GetUserTeams(c echo.Context, userID uint) ([]dto.Member, error) {
	members, err := s.repo.GetUserTeams(userID)
	if err != nil {
		c.Logger().Errorf("Service | UserTeamService | GetTeams (%d): %w", userID, err)
		return nil, tools.ErrInternal("database failure", err.Error())
	}

	dtos := make([]dto.Member, len(members))
	for i, u := range members {
		dtos[i] = dto.Member{
			UserID:   u.UserID,
			Name:     u.User.Name,
			Email:    u.User.Email,
			TeamID:   u.TeamID,
			TeamName: u.Team.Name,
			Role:     u.Role,
		}
	}
	return dtos, nil
}
