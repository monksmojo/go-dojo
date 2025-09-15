package routers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/handlers"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/routers/venue"
)

func RegisterAllRoutes(chiRouter *chi.Mux) {
	projectHandler := handlers.NewProjectHandler()
	venueRoutes := venue.NewVenueRouter(chiRouter, projectHandler)
	venueRoutes.GetVenueRoutes()
	logAllRoutes(chiRouter)
}

func logAllRoutes(chiRouter *chi.Mux) {
	err := chi.Walk(chiRouter, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s", method, route)
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking routes: %v", err)
	}
}
