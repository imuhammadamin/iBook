package main

import (
	"context"
	"iBook/api"
	"iBook/config"
	"iBook/storage/postgres"
	"log"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	defer store.CloseDB()

	c := api.SetupRouter(store)

	log.Println("THE APP IS STARTING...")

	c.Run(":8080")
}
