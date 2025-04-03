package config

import "api/types"

func (p *PostgresStore) CreateUser(user *types.User) error {
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
