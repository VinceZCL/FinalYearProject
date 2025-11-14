package repository

import (
	"scrum.com/internal/client"
	"scrum.com/types/models"
)

type UserRepository interface {
	GetMembers(teamID int) ([]models.UserTeam, error)
}

type userRepository struct {
	client *client.PostgresClient
}

func NewUserRepository(dbclient *client.PostgresClient) UserRepository {
	return &userRepository{client: dbclient}
}

func (r *userRepository) GetMembers(teamID int) ([]models.UserTeam, error) {
	var userTeams []models.UserTeam
	err := r.client.DB.Preload("User").Preload("Team").Where("team_id = ?", teamID).Find(&userTeams).Error
	if err != nil {
		return nil, err
	}
	return userTeams, nil
}
