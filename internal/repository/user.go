package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/models"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
}

type userRepository struct {
	client *client.PostgresClient
}

func NewUserRepository(dbclient *client.PostgresClient) UserRepository {
	return &userRepository{client: dbclient}
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.client.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
