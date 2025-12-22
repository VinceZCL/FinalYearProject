package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/types/model"
	"github.com/VinceZCL/FinalYearProject/types/model/dto"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/labstack/echo/v4"
)

type CheckInService struct {
	repo repository.CheckInRepository
}

func NewCheckInService(repo repository.CheckInRepository) *CheckInService {
	return &CheckInService{repo: repo}
}

func (s *CheckInService) GetUserCheckIns(c echo.Context, userID uint, date string) ([]dto.CheckIn, error) {

	checkIns, err := s.repo.GetUserCheckIns(userID, date)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | GetUserCheckIns: %w", err)
		return nil, err
	}

	dtos := make([]dto.CheckIn, len(checkIns))
	for i, u := range checkIns {
		dtos[i] = dto.CheckIn{
			ID:         u.ID,
			Type:       u.Type,
			Item:       u.Item,
			Jira:       u.Jira,
			Visibility: u.Visibility,
			TeamID:     u.TeamID,
			UserID:     u.UserID,
			Username:   u.User.Name,
			CreatedAt:  u.CreatedAt,
		}
	}
	return dtos, nil
}

func (s *CheckInService) GetTeamCheckIns(c echo.Context, teamID uint, date string) ([]dto.CheckIn, error) {

	checkIns, err := s.repo.GetTeamCheckIns(teamID, date)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | GetTeamCheckIns: %w", err)
		return nil, err
	}

	dtos := make([]dto.CheckIn, len(checkIns))
	for i, u := range checkIns {
		dtos[i] = dto.CheckIn{
			ID:         u.ID,
			Type:       u.Type,
			Item:       u.Item,
			Jira:       u.Jira,
			Visibility: u.Visibility,
			TeamID:     u.TeamID,
			UserID:     u.UserID,
			Username:   u.User.Name,
			CreatedAt:  u.CreatedAt,
		}
	}
	return dtos, nil
}

func (s *CheckInService) GetCheckIn(c echo.Context, checkInID uint) (*dto.CheckIn, error) {
	checkIn, err := s.repo.GetCheckIn(checkInID)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | GetCheckIn (%d): %w", checkInID, err)
		return nil, err
	}

	dto := &dto.CheckIn{
		ID:         checkIn.ID,
		Type:       checkIn.Type,
		Item:       checkIn.Item,
		Jira:       checkIn.Jira,
		Visibility: checkIn.Visibility,
		TeamID:     checkIn.TeamID,
		UserID:     checkIn.UserID,
		Username:   checkIn.User.Name,
		CreatedAt:  checkIn.CreatedAt,
	}
	return dto, nil
}

func (s *CheckInService) NewCheckIn(c echo.Context, req param.NewCheckIn) (*dto.CheckIn, error) {
	input := model.CheckIn{
		UserID:     req.UserID,
		Type:       req.Type,
		Item:       req.Item,
		Jira:       req.Jira,
		Visibility: req.Visibility,
		TeamID:     req.TeamID,
	}

	checkIn, err := s.repo.NewCheckIn(input)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | NewCheckIn: %w", err)
		return nil, err
	}

	dto := &dto.CheckIn{
		ID:         checkIn.ID,
		Type:       checkIn.Type,
		Item:       checkIn.Item,
		Jira:       checkIn.Jira,
		Visibility: checkIn.Visibility,
		TeamID:     checkIn.TeamID,
		UserID:     checkIn.UserID,
		Username:   checkIn.User.Name,
		CreatedAt:  checkIn.CreatedAt,
	}
	return dto, nil
}
