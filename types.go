package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Id           int    `json:"id"`
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func NewUser(username, email, password string) *User {
	passwordHash, err := HashPassword(password)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &User{
		Id:           rand.Intn(100000), // Using it for now to test it
		UserName:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
}

type Event struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	StartTime time.Time   `json:"start_time"`
	Status    EventStatus `json:"status"`
}

func NewEvent(id int, name string, start time.Time, status EventStatus) *Event {
	return &Event{
		Id:        rand.Intn(1000000000000000), // Using it for now to test it
		Name:      name,
		StartTime: start,
		Status:    status,
	}
}

type Bet struct {
	UserId   int       `json:"user_id"`
	EventId  int       `json:"event_id"`
	Amount   float32   `json:"amount"`
	Choice   string    `json:"choice"`
	PlacedAt time.Time `json:"placed_at"`
}

type EventStatus string

const (
	Upcoming  EventStatus = "upcoming"
	Ongoing   EventStatus = "ongoing"
	Completed EventStatus = "completed"
)
