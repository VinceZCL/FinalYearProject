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

func (s *CheckInService) GetUserCheckIns(c echo.Context, userID uint, date string) (*dto.DailyCheckIn, error) {

	checkIns, err := s.repo.GetUserCheckIns(userID, date)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | GetUserCheckIns: %w", err)
		return nil, err
	}

	dailyCheckIn := &dto.DailyCheckIn{
		UserID:    userID,
		Username:  checkIns[0].User.Name, // Assuming all check-ins belong to the same user
		CreatedAt: checkIns[0].CreatedAt, // Assuming created_at is the same for all check-ins
	}

	for _, ci := range checkIns {
		single := &dto.CheckIn{
			ID:         ci.ID,
			Type:       ci.Type,
			Item:       ci.Item,
			Jira:       ci.Jira,
			Visibility: ci.Visibility,
			TeamID:     ci.TeamID,
			UserID:     ci.UserID,
			Username:   ci.User.Name,
			CreatedAt:  ci.CreatedAt,
		}

		switch ci.Type {
		case "yesterday":
			dailyCheckIn.Yesterday = append(dailyCheckIn.Yesterday, single)
		case "today":
			dailyCheckIn.Today = append(dailyCheckIn.Today, single)
		case "blockers":
			dailyCheckIn.Blockers = append(dailyCheckIn.Blockers, single)
		}
	}

	return dailyCheckIn, nil
}

func (s *CheckInService) GetTeamCheckIns(c echo.Context, teamID uint, date string) ([]dto.DailyCheckIn, error) {

	checkIns, err := s.repo.GetTeamCheckIns(teamID, date)
	if err != nil {
		c.Logger().Errorf("Service | CheckInService | GetTeamCheckIns: %w", err)
		return nil, err
	}

	grouped := make(map[uint]*dto.DailyCheckIn)

	for _, ci := range checkIns {
		if _, ok := grouped[ci.UserID]; !ok {
			grouped[ci.UserID] = &dto.DailyCheckIn{
				UserID:    ci.UserID,
				Username:  ci.User.Name,
				CreatedAt: ci.CreatedAt,
			}
		}

		single := &dto.CheckIn{
			ID:         ci.ID,
			Type:       ci.Type,
			Item:       ci.Item,
			Jira:       ci.Jira,
			Visibility: ci.Visibility,
			TeamID:     ci.TeamID,
			UserID:     ci.UserID,
			Username:   ci.User.Name,
			CreatedAt:  ci.CreatedAt,
		}

		switch ci.Type {
		case "yesterday":
			grouped[ci.UserID].Yesterday = append(grouped[ci.UserID].Yesterday, single)
		case "today":
			grouped[ci.UserID].Today = append(grouped[ci.UserID].Today, single)
		case "blockers":
			grouped[ci.UserID].Blockers = append(grouped[ci.UserID].Blockers, single)
		}
	}

	dtos := make([]dto.DailyCheckIn, 0, len(grouped))
	for _, v := range grouped {
		dtos = append(dtos, *v)
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

func (s *CheckInService) BulkCheckIn(c echo.Context, req param.BulkCheckIn) ([]*dto.CheckIn, error) {
	var result []*dto.CheckIn

	// Validate the bulk check-in request
	if err := req.Validate(); err != nil {
		return nil, err // Return error if bulk validation fails
	}

	// Iterate through each NewCheckIn and process it
	for _, checkInReq := range req.CheckIns {
		input := model.CheckIn{
			UserID:     checkInReq.UserID,
			Type:       checkInReq.Type,
			Item:       checkInReq.Item,
			Jira:       checkInReq.Jira,
			Visibility: checkInReq.Visibility,
			TeamID:     checkInReq.TeamID,
		}

		// Call repo to create the check-in in the database
		checkIn, err := s.repo.NewCheckIn(input)
		if err != nil {
			c.Logger().Errorf("Service | CheckInService | NewBulkCheckIn: %w", err)
			continue // Proceed with next check-in even if one fails
		}

		// Map the database model to the DTO
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

		// Add the created check-in DTO to the result slice
		result = append(result, dto)
	}

	return result, nil
}
