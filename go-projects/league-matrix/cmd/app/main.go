package main

import (
	_ "github.com/monksmojo/go-dojo/go-project/league-matrix/docs"
	"github.com/monksmojo/go-dojo/go-project/league-matrix/server"
)

// @title LEAGUE_MATRIX
// @version 1.0
// @description Performing Operations on Matrix
// @host localhost:8080
// @BasePath /
func main() {
	server.Run()
}
