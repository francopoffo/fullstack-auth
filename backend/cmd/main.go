package main

import (
	"log"

	"github.com/francopoffo/fullstack-auth/cmd/api"
)

func main() {

	db, err := api.NewPostgresDB()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		panic(err)
	}
}
