package config

func (p *PostgresStore) CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		user_name VARCHAR(50) NOT NULL, 
		email VARCHAR(50) NOT NULL UNIQUE, 
		PasswordHash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := p.db.Exec(query)

	return err
}

func (p *PostgresStore) CreateEventTable() error {
	query := `CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
    	name VARCHAR(255) NOT NULL,
    	start_time TIMESTAMP NOT NULL,
    	status VARCHAR(50) CHECK (status IN ('upcoming', 'ongoing', 'completed')) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := p.db.Exec(query)
	return err
}

func (p *PostgresStore) CreateBetTable() error {

	query := `CREATE TABLE IF NOT EXISTS bets (
		id SERIAL PRIMARY KEY,
		user_id INT,
		event_id INT,
		amount DECIMAL(10,2) NOT NULL,
		choice VARCHAR(255) NOT NULL,
		placed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (event_id) REFERENCES events(id)
	)`

	_, err := p.db.Exec(query)
	return err
}
