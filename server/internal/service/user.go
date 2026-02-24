package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model"
	"github.com/VinceZCL/FinalYearProject/types/model/dto"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
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
		return nil, tools.ErrInternal("database failure", err.Error())
	}

	dtos := make([]dto.User, len(users))
	for i, u := range users {
		dtos[i] = dto.User{
			ID:     u.ID,
			Name:   u.Name,
			Email:  u.Email,
			Type:   u.Type,
			Status: u.Status,
		}
	}
	return dtos, nil
}

func (s *UserService) GetUser(c echo.Context, userID uint) (*dto.User, error) {
	user, err := s.repo.GetUser(userID)
	if err != nil {
		c.Logger().Errorf("Service | UserService | GetUser (%d): %w", userID, err)
		return nil, tools.ErrInternal("database failure", err.Error())
	}

	dto := &dto.User{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Type:   user.Type,
		Status: user.Status,
	}
	return dto, nil
}

// now used in Admin user creation
func (s *UserService) NewUser(c echo.Context, req param.NewUser) (*dto.User, error) {
	var userType string
	if req.Type != "" {
		userType = req.Type
	} else {
		userType = "user"
	}
	hashed, err := tools.HashPass(req.Password)
	if err != nil {
		return nil, tools.ErrInternal("hasing error", err.Error())
	}
	input := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Type:     userType,
		Status:   "active",
	}

	user, err := s.repo.NewUser(input)
	if err != nil {
		c.Logger().Errorf("Service | UserService | NewUser: %w", err)
		return nil, tools.ErrInternal("database failure", err.Error())
	}

	dto := &dto.User{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Type:   user.Type,
		Status: user.Status,
	}
	return dto, nil
}

func (s *UserService) DeactivateUser(c echo.Context, userID uint) (*dto.User, error) {

	user, err := s.repo.GetUser(userID)
	if err != nil {
		c.Logger().Errorf("Service | UserService | DeactivateUser: %w", err)
		return nil, tools.ErrNotFound("database failure", "User not found")
	}

	if user.Status == "active" {
		user, err = s.repo.DeactivateUser(*user)
		if err != nil {
			c.Logger().Errorf("Service | UserService | DeactivateUser: %w", err)
			return nil, tools.ErrInternal("database failure", err.Error())
		}
	} else {
		user, err = s.repo.ActivateUser(*user)
		if err != nil {
			c.Logger().Errorf("Service | UserService | ActivateUser: %w", err)
			return nil, tools.ErrInternal("database failure", err.Error())
		}
	}

	dto := &dto.User{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Type:   user.Type,
		Status: user.Status,
	}

	return dto, nil
}
