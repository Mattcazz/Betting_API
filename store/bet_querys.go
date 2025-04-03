package store

import (
	"api/models"
	"database/sql"
	"fmt"
)

func (p *PostgresStore) CreateBet(bet *models.Bet) error {
	query := `INSERT INTO bets 
			(user_id, event_id, amount, choice, placed_at)
			VALUES ($1, $2, $3, $4, $5)`

	_, err := p.db.Query(query,
		bet.UserId,
		bet.EventId,
		bet.Amount,
		bet.Choice,
		bet.PlacedAt)

	return err
}

func (p *PostgresStore) GetBets() ([]*models.Bet, error) {
	rows, err := p.db.Query("SELECT * FROM bets")

	bets := []*models.Bet{}

	for rows.Next() {
		bet, err := scanBetRow(rows)
		if err != nil {
			return nil, err
		}
		bets = append(bets, bet)
	}

	return bets, err
}

func (p *PostgresStore) GetBetById(id int) (*models.Bet, error) {
	row, err := p.db.Query("SELECT * FROM bets WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		return scanBetRow(row)
	}

	return nil, fmt.Errorf("the query came up with no results")
}

func (p *PostgresStore) DeleteBetById(id int) error {
	_, err := p.db.Query("DELETE FROM bets WHERE id = $1", id)

	return err
}

// Private function that returns a bet given a row to scan
func scanBetRow(row *sql.Rows) (*models.Bet, error) {
	bet := new(models.Bet)

	err := row.Scan(
		&bet.UserId,
		&bet.EventId,
		&bet.Amount,
		&bet.Choice,
		&bet.PlacedAt)
	return bet, err
}
