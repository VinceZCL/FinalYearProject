package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type UserTeamRepository interface {
	GetMembers(teamID uint) ([]model.UserTeam, error)
	NewMember(input model.UserTeam) (*model.UserTeam, error)
	GetUserTeams(userID uint) ([]model.UserTeam, error)
	IsTeamAdmin(userID uint, teamID uint) (bool, error)
	DeleteMember(teamID uint, userID uint) error
}

type userTeamRepository struct {
	client *client.PostgresClient
}

func NewUserTeamRepository(dbclient *client.PostgresClient) UserTeamRepository {
	return &userTeamRepository{client: dbclient}
}

func (r *userTeamRepository) GetMembers(teamID uint) ([]model.UserTeam, error) {
	var userTeams []model.UserTeam
	err := r.client.DB.Joins("JOIN fyp_scrum_users ON fyp_scrum_users.id = fyp_scrum_user_teams.user_id AND fyp_scrum_users.status = ?", "active").Preload("User").Preload("Team").Where("fyp_scrum_user_teams.team_id = ?", teamID).Find(&userTeams).Error
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
	if err := r.client.DB.Joins("JOIN fyp_scrum_users ON fyp_scrum_users.id = fyp_scrum_user_teams.user_id AND fyp_scrum_users.status = ?", "active").Preload("User").Preload("Team").Where("fyp_scrum_user_teams.user_id = ? AND fyp_scrum_user_teams.team_id = ?", input.UserID, input.TeamID).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userTeamRepository) GetUserTeams(userID uint) ([]model.UserTeam, error) {
	var userTeams []model.UserTeam
	err := r.client.DB.Joins("JOIN fyp_scrum_users ON fyp_scrum_users.id = fyp_scrum_user_teams.user_id AND fyp_scrum_users.status = ?", "active").Preload("User").Preload("Team").Where("fyp_scrum_user_teams.user_id = ?", userID).Find(&userTeams).Error
	if err != nil {
		return nil, err
	}
	return userTeams, nil
}

func (r *userTeamRepository) IsTeamAdmin(userID uint, teamID uint) (bool, error) {
	var exists int
	err := r.client.DB.Model(&model.UserTeam{}).Select("1").Where("fyp_scrum_user_teams.user_id = ? AND fyp_scrum_user_teams.team_id = ? AND fyp_scrum_user_teams.role = ?", userID, teamID, "admin").Limit(1).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *userTeamRepository) DeleteMember(teamID uint, userID uint) error {
	err := r.client.DB.Model(&model.UserTeam{}).Where("team_id = ? AND user_id = ?", teamID, userID).Delete(&model.UserTeam{}).Error
	if err != nil {
		return err
	}
	return nil
}
