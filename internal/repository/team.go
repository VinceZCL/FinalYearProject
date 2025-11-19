package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/models"
)

type TeamRepository interface {
	GetTeams() ([]models.Team, error)
}

type teamRepository struct {
	client *client.PostgresClient
}

func NewTeamRepository(dbclient *client.PostgresClient) TeamRepository {
	return &teamRepository{client: dbclient}
}

func (r *teamRepository) GetTeams() ([]models.Team, error) {
	var teams []models.Team
	err := r.client.DB.Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}
