package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type CheckInRepository interface {
	GetUserCheckIns(userID uint) ([]model.CheckIn, error)
	GetTeamCheckIns(teamID uint) ([]model.CheckIn, error)
	GetCheckIn(checkInID uint) (*model.CheckIn, error)
	NewCheckIn(input model.CheckIn) (*model.CheckIn, error)
}

type checkInRepository struct {
	client *client.PostgresClient
}

func NewCheckInRepository(dbclient *client.PostgresClient) CheckInRepository {
	return &checkInRepository{client: dbclient}
}

func (r *checkInRepository) GetUserCheckIns(userID uint) ([]model.CheckIn, error) {
	var checkIns []model.CheckIn
	err := r.client.DB.Preload("User").Preload("Team").Where("user_id = ?", userID).Find(&checkIns).Error
	if err != nil {
		return nil, err
	}
	return checkIns, nil
}

func (r *checkInRepository) GetTeamCheckIns(teamID uint) ([]model.CheckIn, error) {
	var checkIns []model.CheckIn
	err := r.client.DB.Preload("User").Preload("Team").Where("team_id = ?", teamID).Find(&checkIns).Error
	if err != nil {
		return nil, err
	}
	return checkIns, nil
}

func (r *checkInRepository) GetCheckIn(checkInID uint) (*model.CheckIn, error) {
	var checkIn *model.CheckIn
	err := r.client.DB.Preload("User").Preload("Team").First(&checkIn, checkInID).Error
	if err != nil {
		return nil, err
	}
	return checkIn, nil
}

func (r *checkInRepository) NewCheckIn(input model.CheckIn) (*model.CheckIn, error) {
	err := r.client.DB.Create(&input).Error
	if err != nil {
		return nil, err
	}
	if err := r.client.DB.Preload("User").Preload("Team").First(&input, input.ID).Error; err != nil {
		return nil, err
	}
	return &input, nil
}
