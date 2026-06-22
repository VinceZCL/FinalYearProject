package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type CheckInRepository interface {
	GetUserCheckIns(userID uint, date string) ([]model.CheckIn, error)
	GetTeamCheckIns(teamID uint, date string) ([]model.CheckIn, error)
	GetCheckIn(checkInID uint) (*model.CheckIn, error)
	NewCheckIn(input model.CheckIn) (*model.CheckIn, error)
	GetYesterday(userID uint) ([]model.CheckIn, error)
	DeleteCheckIns(userID uint) error
}

type checkInRepository struct {
	client *client.PostgresClient
}

func NewCheckInRepository(dbclient *client.PostgresClient) CheckInRepository {
	return &checkInRepository{client: dbclient}
}

func (r *checkInRepository) GetUserCheckIns(userID uint, date string) ([]model.CheckIn, error) {
	var checkIns []model.CheckIn
	start, end, err := tools.GetTimes(date)
	if err != nil {
		return nil, err
	}
	err = r.client.DB.Preload("User").Preload("Team").
		Where("fyp_scrum_checkins.user_id = ?", userID).
		Where("fyp_scrum_checkins.created_at >= ? AND fyp_scrum_checkins.created_at < ?", start, end).
		Find(&checkIns).Error
	if err != nil {
		return nil, err
	}
	return checkIns, nil
}

func (r *checkInRepository) GetTeamCheckIns(teamID uint, date string) ([]model.CheckIn, error) {
	var checkIns []model.CheckIn
	start, end, err := tools.GetTimes(date)
	if err != nil {
		return nil, err
	}

	err = r.client.DB.Model(&model.CheckIn{}).
		Where(`
			EXISTS (
				SELECT 1
				FROM fyp_scrum_user_teams ut
				WHERE ut.user_id = fyp_scrum_checkins.user_id
				AND ut.team_id = ?
			)
		`, teamID).
		Where(`fyp_scrum_checkins.created_at >= ? AND fyp_scrum_checkins.created_at < ?`, start, end).
		Where(`
			(
				fyp_scrum_checkins.visibility = 'all'
				OR (
					fyp_scrum_checkins.visibility = 'team'
					AND fyp_scrum_checkins.team_id = ?	
				)
			)
		`, teamID).
		Preload("User").
		Preload("Team").
		Preload("Comments", "fyp_scrum_comments.team_id = ?", teamID).
		Preload("Comments.User").
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

func (r *checkInRepository) GetYesterday(userID uint) ([]model.CheckIn, error) {
	var checkIn []model.CheckIn

	sub := r.client.DB.Model(&model.CheckIn{}).
		Select("DATE(fyp_scrum_checkins.created_at)").
		Where("fyp_scrum_checkins.user_id = ?", userID).
		Order("fyp_scrum_checkins.created_at DESC").
		Limit(1)

	err := r.client.DB.Preload("User").Preload("Team").
		Where("fyp_scrum_checkins.user_id = ?", userID).
		Where("DATE(fyp_scrum_checkins.created_at) = (?)", sub).
		Order("fyp_scrum_checkins.created_at DESC").
		Find(&checkIn).Error

	if err != nil {
		return nil, err
	}

	return checkIn, nil
}

func (r *checkInRepository) DeleteCheckIns(userID uint) error {

	start, end, err := tools.GetTimes("")
	if err != nil {
		return err
	}

	err = r.client.DB.
		Where("fyp_scrum_checkins.user_id = ?", userID).
		Where("fyp_scrum_checkins.created_at >= ? AND fyp_scrum_checkins.created_at < ?", start, end).
		Delete(&model.CheckIn{}).Error

	if err != nil {
		return err
	}

	return nil
}
