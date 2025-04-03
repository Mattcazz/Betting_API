package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("HOST")
	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASE")
	dbPort := os.Getenv("PORT")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName) //postgres://USER:PASSWORD@HOST:PORT/DATABASE?OPTIONS

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(10 * time.Minute) // Every connection that opens is going to close after 10 mins

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	fmt.Printf("Connected to the %s database successfully!\n", dbName)

	return &PostgresStore{
		db: db,
	}, nil

}

func (p *PostgresStore) Init() error {
	if err := p.CreateEventTable(); err != nil {
		return err
	}
	if err := p.CreateUserTable(); err != nil {
		return err
	}

	if err := p.CreateBetTable(); err != nil {
		return err
	}

	return nil
}

// CloseDB closes the database connection
func (p *PostgresStore) CloseDB() {
	if p.db != nil {
		p.db.Close()
		fmt.Println("Database closed!!")
	}
}
