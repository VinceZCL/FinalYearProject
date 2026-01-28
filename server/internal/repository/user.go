package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUser(userID uint) (*model.User, error)
	NewUser(model.User) (*model.User, error)
	DeactivateUser(user model.User) (*model.User, error)
	ActivateUser(user model.User) (*model.User, error)
}

type userRepository struct {
	client *client.PostgresClient
}

func NewUserRepository(dbclient *client.PostgresClient) UserRepository {
	return &userRepository{client: dbclient}
}

func (r *userRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	err := r.client.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUser(userID uint) (*model.User, error) {
	var user model.User
	err := r.client.DB.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) NewUser(user model.User) (*model.User, error) {
	err := r.client.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) DeactivateUser(user model.User) (*model.User, error) {

	user.Status = "deactivated"
	err := r.client.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) ActivateUser(user model.User) (*model.User, error) {

	user.Status = "active"
	err := r.client.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
