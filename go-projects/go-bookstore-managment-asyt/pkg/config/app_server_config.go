package config

import (
	"log"
	"net/http"

	"github.com/monks_mojo/go-dojo/go-projects/go-bookstore-managment-asyt/pkg/routes"
)

type apiServer struct {
	port     string
	serveMux *http.ServeMux
}

func (apiServer *apiServer) Run() error {
	loadEnv()
	connectToDB()
	bootStrapDB()
	routes.RegisterBookRoutes(apiServer.serveMux)
	server := http.Server{
		Addr:    apiServer.port,
		Handler: apiServer.serveMux,
	}
	log.Printf("starting the server at port %v", apiServer.port)
	return server.ListenAndServe()
}

func NewApiServer(addr string) *apiServer {
	return &apiServer{
		port:     addr,
		serveMux: http.NewServeMux(),
	}
}
