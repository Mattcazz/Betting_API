package main

import (
	"api/server"
	"api/store"
	"fmt"
)

func main() {
	pgStore, err := store.NewPostgresStore()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer pgStore.CloseDB()

	if err := pgStore.Init(); err != nil {
		fmt.Println(err.Error())
	}

	apiServer := server.NewApiServer(pgStore)
	apiServer.SetUpRoutes()
	apiServer.Engine.Run("localhost:8080")
	fmt.Println("Running server on port 8080")
}
