package store

import "api/models"

type Store interface {
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	DeleteUser(id int) error
}
