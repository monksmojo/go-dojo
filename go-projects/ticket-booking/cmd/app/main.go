package main

import (
	"log"

	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/config"
)

func main() {
	httpServer := config.NewServer(":8080")
	if err := httpServer.Run(); err != nil {
		log.Fatal("difficulty in starting the server", err)
	}
	log.Printf("server started st port 8080")
}
