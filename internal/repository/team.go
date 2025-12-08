package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type TeamRepository interface {
	GetTeams() ([]model.Team, error)
	GetTeam(teamID uint) (*model.Team, error)
	NewTeam(team model.Team) (*model.Team, error)
}

type teamRepository struct {
	client *client.PostgresClient
}

func NewTeamRepository(dbclient *client.PostgresClient) TeamRepository {
	return &teamRepository{client: dbclient}
}

func (r *teamRepository) GetTeams() ([]model.Team, error) {
	var teams []model.Team
	err := r.client.DB.Preload("Creator").Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *teamRepository) GetTeam(teamID uint) (*model.Team, error) {
	var team *model.Team
	err := r.client.DB.Preload("Creator").First(&team, teamID).Error
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (r *teamRepository) NewTeam(team model.Team) (*model.Team, error) {
	err := r.client.DB.Create(&team).Error
	if err != nil {
		return nil, err
	}
	if err := r.client.DB.Preload("Creator").First(&team, team.ID).Error; err != nil {
		return nil, err
	}
	return &team, err
}
