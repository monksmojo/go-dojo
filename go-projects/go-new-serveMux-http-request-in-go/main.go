package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("How to create your Golang APIs from now on with v1.22")
	httpServer := NewApiServer("8080")
	if err := httpServer.Run(); err != nil {
		log.Fatal("difficulty in starting the server", err)
	}
}
