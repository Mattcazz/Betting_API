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

func (p *PostgresStore) GetUsers() ([]*types.User, error) {
	rows, err := p.db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	users := []*types.User{}

	for rows.Next() {
		user := new(types.User)
		err := rows.Scan(
			&user.Id,
			&user.UserName,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *PostgresStore) GetUserById(id int) (*types.User, error) {
	query := `SELECT * FROM USERS WHERE id = $1`

	row, err := p.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	user := new(types.User)

	row.Next()
	err2 := row.Scan(
		&user.Id,
		&user.UserName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt)

	if err2 != nil {
		return nil, err
	}

	return user, nil
}
