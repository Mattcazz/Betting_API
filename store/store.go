package store

import "api/models"

type Store interface {
	// User functions
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	DeleteUserById(id int) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserBets(id int) ([]*models.Bet, error)

	// Event functions
	CreateEvent(event *models.Event) error
	GetEvents() ([]*models.Event, error)
	GetEventById(id int) (*models.Event, error)
	DeleteEventById(id int) error

	// Bet funtions
	CreateBet(bet *models.Bet) error
	GetBets() ([]*models.Bet, error)
	GetBet(user_id, event_id int) (*models.Bet, error)
	DeleteBet(user_id, event_id int) error
}
