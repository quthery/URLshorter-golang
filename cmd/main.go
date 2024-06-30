package main

import (
	"log"
	"shorter/pkg/handler"
	"shorter/pkg/server"
	sqlite "shorter/pkg/sqlite"
)

func main() {
	storage, err := sqlite.New("database/storage.db")

	if err != nil {
		log.Fatalf("db error: %s", err)
	}

	handler := handler.NewHandler(storage)
	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error pri zapyske: %s", err)
	}
}
