package main

import "github.com/francopoffo/fullstack-auth/cmd/api"

func main() {

	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		panic(err)
	}
}
