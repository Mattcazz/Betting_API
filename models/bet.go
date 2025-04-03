package models

import "time"

type Bet struct {
	UserId   int       `json:"user_id"`
	EventId  int       `json:"event_id"`
	Amount   float32   `json:"amount"`
	Choice   string    `json:"choice"`
	PlacedAt time.Time `json:"placed_at"`
}

type NewBetRequest struct {
	Choice string  `json:"choice"`
	Amount float32 `json:"amount"`
}

func NewBet(userId, eventId int, amount float32, choice string) *Bet {
	return &Bet{
		UserId:   userId,
		EventId:  eventId,
		Amount:   amount,
		Choice:   choice,
		PlacedAt: time.Now(),
	}
}
