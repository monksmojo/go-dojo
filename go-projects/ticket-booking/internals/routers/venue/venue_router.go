package venue

import (
	"github.com/go-chi/chi/v5"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/handlers"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/handlers/venue"
)

func NewVenueRouter(chiRouter *chi.Mux, projectHandler handlers.ProjectHandler) *VenueRouter {
	return &VenueRouter{
		chiRouter:    chiRouter,
		venueHandler: projectHandler.GetVenueHandler(),
	}
}

type VenueRouter struct {
	chiRouter    *chi.Mux
	venueHandler *venue.VenueHandler
}

func (vr *VenueRouter) GetVenueRoutes() {
	vr.chiRouter.Get("/venue-homepage", vr.venueHandler.GetVenueHomepage)
	vr.chiRouter.Get("/venue-registration-form", vr.venueHandler.GetVenueRegistrationForm)
	vr.chiRouter.Post("/venue", vr.venueHandler.CreateVenue)
	vr.chiRouter.Get("/venue-names",vr.venueHandler.GetAllVenueNames)
}
