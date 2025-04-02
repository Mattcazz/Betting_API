package db

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

var db *sql.DB

func Connect() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("HOST")
	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASE")
	dbPort := os.Getenv("PORT")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName) //postgres://USER:PASSWORD@HOST:PORT/DATABASE?OPTIONS

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(10 * time.Minute) // Every connection that opens is going to close after 10 mins

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return err
	}

	fmt.Printf("Connected to the %s database successfully!\n", dbName)

	return nil

}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Database closed!!")
	}
}

func GetDB() *sql.DB {
	return db
}
