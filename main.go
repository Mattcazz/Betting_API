package main

import (
	"api/config"
	"fmt"
)

var pStore *config.PostgresStore

func main() {

	var err error
	pStore, err = config.NewPostgresStore()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer pStore.CloseDB()

	if err := pStore.Init(); err != nil {
		fmt.Println(err.Error())
	}

	apiServer := NewApiServer()
	apiServer.SetUpRoutes()
	apiServer.engine.Run("localhost:8080")
	fmt.Println("Running server on port 8080")
}
