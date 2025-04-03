package types

import "time"

type Event struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	StartTime time.Time   `json:"start_time"`
	Status    EventStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

func NewEvent(id int, name string, start time.Time, status EventStatus) *Event {
	return &Event{
		Name:      name,
		StartTime: start,
		Status:    status,
	}
}

type EventStatus string

const (
	Upcoming  EventStatus = "upcoming"
	Ongoing   EventStatus = "ongoing"
	Completed EventStatus = "completed"
)
