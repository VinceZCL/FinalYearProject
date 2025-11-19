package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/models"
)

type UserTeamRepository interface {
	GetMembers(teamID int) ([]models.UserTeam, error)
}

type userTeamRepository struct {
	client *client.PostgresClient
}

func NewUserTeamRepository(dbclient *client.PostgresClient) UserTeamRepository {
	return &userTeamRepository{client: dbclient}
}

func (r *userTeamRepository) GetMembers(teamID int) ([]models.UserTeam, error) {
	var userTeams []models.UserTeam
	err := r.client.DB.Preload("User").Preload("Team").Where("team_id = ?", teamID).Find(&userTeams).Error
	if err != nil {
		return nil, err
	}
	return userTeams, nil
}
