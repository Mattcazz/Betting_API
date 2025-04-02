package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *gorm.DB

func Init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Env file sucks!")
	}

	dbHost := os.Getenv("HOST")
	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	db, err = postgres.Open(dsn)

	if err != nil {
		fmt.Println("Not correctly connected")
	} else {
		fmt.Println("Connected to the DataBase called", dbName)
	}
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close the database connection:", err)
	}
	sqlDB.Close()
}
