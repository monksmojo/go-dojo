package config

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/db"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/routers"
)

type server struct {
	port      string
	chiRouter *chi.Mux
}

func NewServer(addr string) *server {
	return &server{
		port:      addr,
		chiRouter: chi.NewRouter(),
	}
}

func (server *server) Run() error {
	loadEnv()
	db.Init()
	db.BootStrapDB()
	routers.RegisterAllRoutes(server.chiRouter)
	httpServer := http.Server{
		Addr:    server.port,
		Handler: server.chiRouter,
	}
	log.Printf("starting server at the port %v", server.port)
	return httpServer.ListenAndServe()
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}
}


