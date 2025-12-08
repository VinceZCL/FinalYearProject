package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type UserTeamRepository interface {
	GetMembers(teamID uint) ([]model.UserTeam, error)
	NewMember(input model.UserTeam) (*model.UserTeam, error)
	GetUserTeams(userID uint) ([]model.UserTeam, error)
}

type userTeamRepository struct {
	client *client.PostgresClient
}

func NewUserTeamRepository(dbclient *client.PostgresClient) UserTeamRepository {
	return &userTeamRepository{client: dbclient}
}

func (r *userTeamRepository) GetMembers(teamID uint) ([]model.UserTeam, error) {
	var userTeams []model.UserTeam
	err := r.client.DB.Preload("User").Preload("Team").Where("team_id = ?", teamID).Find(&userTeams).Error
	if err != nil {
		return nil, err
	}
	return userTeams, nil
}

func (r *userTeamRepository) NewMember(input model.UserTeam) (*model.UserTeam, error) {
	err := r.client.DB.Create(&input).Error
	if err != nil {
		return nil, err
	}
	var result model.UserTeam
	if err := r.client.DB.Preload("User").Preload("Team").Where("user_id = ? AND team_id = ?", input.UserID, input.TeamID).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userTeamRepository) GetUserTeams(userID uint) ([]model.UserTeam, error) {
	var userTeams []model.UserTeam
	err := r.client.DB.Preload("User").Preload("Team").Where("user_id = ?", userID).Find(&userTeams).Error
	if err != nil {
		return nil, err
	}
	return userTeams, nil
}
