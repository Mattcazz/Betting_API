package store

import "api/models"

type Store interface {
	// User functions
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	DeleteUserById(id int) error

	// Event functions
	CreateEvent(event *models.Event) error
	GetEvents() ([]*models.Event, error)
	GetEventById(id int) (*models.Event, error)
	DeleteEventById(id int) error

	// Bet funtions
	CreateBet(bet *models.Bet) error
	GetBets() ([]*models.Bet, error)
	GetBetById(id int) (*models.Bet, error)
	DeleteBetById(id int) error
}
