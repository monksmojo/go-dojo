package main

import (
	"log"

	"github.com/monks_mojo/go-dojo/go-projects/go-bookstore-managment-asyt/pkg/config"
)

func main() {
	httpServer := config.NewApiServer(":8080")
	if err := httpServer.Run(); err != nil {
		log.Fatal("difficulty in starting the server", err)
	}
}
