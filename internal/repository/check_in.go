package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/models"
)

type CheckInRepository interface {
	GetUserCheckIns(userID int) ([]models.CheckIn, error)
	GetTeamCheckIns(teamID int) ([]models.CheckIn, error)
}

type checkInRepository struct {
	client *client.PostgresClient
}

func NewCheckInRepository(dbclient *client.PostgresClient) CheckInRepository {
	return &checkInRepository{client: dbclient}
}

func (r *checkInRepository) GetUserCheckIns(userID int) ([]models.CheckIn, error) {
	var checkIns []models.CheckIn
	err := r.client.DB.Preload("User").Preload("Team").Where("user_id = ?", userID).Find(&checkIns).Error
	if err != nil {
		return nil, err
	}
	return checkIns, nil
}

func (r *checkInRepository) GetTeamCheckIns(teamID int) ([]models.CheckIn, error) {
	var checkIns []models.CheckIn
	err := r.client.DB.Preload("User").Preload("Team").Where("team_id = ?", teamID).Find(&checkIns).Error
	if err != nil {
		return nil, err
	}
	return checkIns, nil
}
