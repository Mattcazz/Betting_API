package store

import "api/types"

type Store interface {
	CreateUser(user *types.User) error
	GetUsers() ([]*types.User, error)
	GetUserById(id int) (*types.User, error)
	DeleteUser(id int) error
}
