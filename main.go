package main

import (
	"api/config"
	"fmt"
)

func main() {
	pStore, err := config.NewPostgresStore()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer pStore.CloseDB()

	if err := pStore.Init(); err != nil {
		fmt.Println(err.Error())
	}
}
