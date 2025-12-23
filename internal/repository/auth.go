package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type AuthRepository interface {
	GetCredentials(email string) (*model.User, error)
}

type authRepository struct {
	client *client.PostgresClient
}

func NewAuthRepository(dbclient *client.PostgresClient) AuthRepository {
	return &authRepository{client: dbclient}
}

func (r *authRepository) GetCredentials(email string) (*model.User, error) {
	var user *model.User
	err := r.client.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
