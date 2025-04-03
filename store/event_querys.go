package store

import (
	"api/models"
	"database/sql"
	"fmt"
)

func (p *PostgresStore) CreateEvent(event *models.Event) error {
	query := `INSERT INTO events  
			(name, start_time, status, created_at)
			VALUES ($1, $2, $3, $4)`

	_, err := p.db.Query(query,
		event.Name,
		event.StartTime,
		event.Status,
		event.CreatedAt)

	return err
}

func (p *PostgresStore) GetEvents() ([]*models.Event, error) {
	rows, err := p.db.Query("SELECT * FROM events")

	events := []*models.Event{}

	for rows.Next() {
		event, err := scanEventRow(rows)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, err

}

func (p *PostgresStore) GetEventById(id int) (*models.Event, error) {
	row, err := p.db.Query("SELECT * FROM events WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		return scanEventRow(row)
	}

	return nil, fmt.Errorf("the query came up with no results")
}

func (p *PostgresStore) DeleteEventById(id int) error {
	_, err := p.db.Query("DELETE FROM events WHERE id = $1", id)

	return err
}

// Private function that returns a event given a row to scan
func scanEventRow(row *sql.Rows) (*models.Event, error) {
	event := new(models.Event)

	err := row.Scan(
		&event.Id,
		&event.Name,
		&event.StartTime,
		&event.Status,
		&event.CreatedAt)

	return event, err
}
