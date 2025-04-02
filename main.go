package main

import (
	"api/db"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func main() {
	if err := db.Connect(); err != nil {
		fmt.Println(err.Error())
	}
	defer db.CloseDB()
}
