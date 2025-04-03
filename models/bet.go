package models

import "time"

type Bet struct {
	UserId   int       `json:"user_id"`
	EventId  int       `json:"event_id"`
	Amount   float32   `json:"amount"`
	Choice   string    `json:"choice"`
	PlacedAt time.Time `json:"placed_at"`
}
