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

func (s *UserService) GetUsers(c echo.Context) ([]dto.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		c.Logger().Errorf("Service | UserService | GetUsers: %w", err)
		return nil, err
	}

	dtos := make([]dto.User, len(users))
	for i, u := range users {
		dtos[i] = dto.User{
			UserID: u.ID,
			Name:   u.Name,
			Email:  u.Email,
			Type:   u.Type,
		}
	}
	return dtos, nil
}

func (s *UserService) GetUser(c echo.Context, userID int) (*dto.User, error) {
	user, err := s.repo.GetUser(userID)
	if err != nil {
		c.Logger().Errorf("Service | UserService | GetUser (%d): %w", userID, err)
		return nil, err
	}

	dto := &dto.User{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Type:   user.Type,
	}
	return dto, nil
}
