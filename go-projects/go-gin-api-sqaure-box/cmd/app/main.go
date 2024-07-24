package main

import (
	_ "github.com/monksmojo/go-dojo/go-project/go-gin-api-square-box/docs"
	"github.com/monksmojo/go-dojo/go-project/go-gin-api-square-box/server"
)

// @title Square-box
// @version 1.0
// @description Performing Operations on Matrix
// @host localhost:8080
// @BasePath /
func main() {
	server.Run()
}
