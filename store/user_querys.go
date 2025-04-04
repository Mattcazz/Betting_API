package store

import (
	"api/models"
	"database/sql"
	"fmt"
)

func (p *PostgresStore) CreateUser(user *models.User) error {
	query := `INSERT INTO users 
			(user_name, email, passwordhash, created_at)
			VALUES ($1, $2, $3, $4)`

	_, err := p.db.Query(query,
		user.UserName,
		user.Email,
		user.PasswordHash,
		user.CreatedAt)

	return err
}

func (p *PostgresStore) GetUsers() ([]*models.User, error) {
	rows, err := p.db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	users := []*models.User{}

	for rows.Next() {

		user, err := scanUserRow(rows)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *PostgresStore) GetUserById(id int) (*models.User, error) {
	query := `SELECT * FROM USERS WHERE id = $1`

	row, err := p.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		return scanUserRow(row)
	}

	return nil, fmt.Errorf("the query came up with no results")
}

func (p *PostgresStore) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT * FROM USERS WHERE email = $1`

	row, err := p.db.Query(query, email)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		return scanUserRow(row)
	}

	return nil, fmt.Errorf("the query came up with no results")
}

func (p *PostgresStore) DeleteUserById(id int) error {
	_, err := p.db.Query("DELETE FROM users WHERE id = $1", id)

	return err
}

// Private function that returns a user given a row to scan
func scanUserRow(row *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := row.Scan(
		&user.Id,
		&user.UserName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt)
	return user, err
}
