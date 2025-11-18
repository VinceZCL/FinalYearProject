package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/types/models/dto"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetMembers(c echo.Context, teamID int) ([]dto.Member, error) {

	users, err := s.repo.GetMembers(teamID)
	if err != nil {
		c.Logger().Errorf("Service | UserService | GetMembers: %v", err)
		return nil, err
	}

	dtos := make([]dto.Member, len(users))
	for i, u := range users {
		dtos[i] = dto.Member{
			Name:     u.User.Name,
			Email:    u.User.Email,
			TeamName: u.Team.Name,
			Role:     u.Role,
		}
	}
	return dtos, nil
}
