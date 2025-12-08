package service

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/types/models/dto"
	"github.com/labstack/echo/v4"
)

type CheckInService struct {
	repo repository.CheckInRepository
}

func NewCheckInService(repo repository.CheckInRepository) *CheckInService {
	return &CheckInService{repo: repo}
}

func (s *CheckInService) GetUserCheckIns(c echo.Context, userID int) ([]dto.CheckIn, error) {

	checkIns, err := s.repo.GetUserCheckIns(userID)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | GetUserCheckIns: %w", err)
		return nil, err
	}

	dtos := make([]dto.CheckIn, len(checkIns))
	for i, u := range checkIns {
		dtos[i] = dto.CheckIn{
			CheckInID:  u.ID,
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

func (s *CheckInService) GetTeamCheckIns(c echo.Context, teamID int) ([]dto.CheckIn, error) {

	checkIns, err := s.repo.GetUserCheckIns(teamID)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | GetTeamCheckIns: %w", err)
		return nil, err
	}

	dtos := make([]dto.CheckIn, len(checkIns))
	for i, u := range checkIns {
		dtos[i] = dto.CheckIn{
			CheckInID:  u.ID,
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

func (s *CheckInService) GetCheckIn(c echo.Context, checkInID int) (*dto.CheckIn, error) {
	checkIn, err := s.repo.GetCheckIn(checkInID)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | GetCheckIn (%d): %w", checkInID, err)
		return nil, err
	}

	dto := &dto.CheckIn{
		CheckInID:  checkIn.ID,
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
