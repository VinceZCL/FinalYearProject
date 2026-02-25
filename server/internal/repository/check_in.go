package repository

import (
	"time"

	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type CheckInRepository interface {
	GetUserCheckIns(userID uint, date string) ([]model.CheckIn, error)
	GetTeamCheckIns(teamID uint, date string) ([]model.CheckIn, error)
	GetCheckIn(checkInID uint) (*model.CheckIn, error)
	NewCheckIn(input model.CheckIn) (*model.CheckIn, error)
}

type checkInRepository struct {
	client *client.PostgresClient
}

func NewCheckInRepository(dbclient *client.PostgresClient) CheckInRepository {
	return &checkInRepository{client: dbclient}
}

func (r *checkInRepository) GetUserCheckIns(userID uint, date string) ([]model.CheckIn, error) {
	var checkIns []model.CheckIn
	start, end, err := getTimes(date)
	if err != nil {
		return nil, err
	}
	err = r.client.DB.Preload("User").Preload("Team").
		Where("user_id = ?", userID).
		Where("created_at >= ? AND created_at < ?", start, end).
		Find(&checkIns).Error
	if err != nil {
		return nil, err
	}
	return checkIns, nil
}

func (r *checkInRepository) GetTeamCheckIns(teamID uint, date string) ([]model.CheckIn, error) {
	var checkIns []model.CheckIn
	start, end, err := getTimes(date)
	if err != nil {
		return nil, err
	}

	err = r.client.DB.Model(&model.CheckIn{}).
		Where(`
			EXISTS (
				SELECT 1
				FROM user_teams ut
				WHERE ut.user_id = check_ins.user_id
				AND ut.team_id = ?
			)
		`, teamID).
		Where(`check_ins.created_at >= ? AND check_ins.created_at < ?`, start, end).
		Where(`
			(
				check_ins.visibility = 'all'
				OR (
					check_ins.visibility = 'team'
					AND check_ins.team_id = ?	
				)
			)
		`, teamID).
		Preload("User").
		Preload("Team").
		Find(&checkIns).Error

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

func getTimes(dateStr string) (start, end time.Time, err error) {

	var day time.Time
	if dateStr == "" {
		day = time.Now()
	} else {
		day, err = time.ParseInLocation(time.DateOnly, dateStr, time.Local)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
	}
	start = time.Date(
		day.Year(),
		day.Month(),
		day.Day(),
		0, 0, 0, 0,
		day.Location(),
	)
	end = start.Add(24 * time.Hour)
	return
}
