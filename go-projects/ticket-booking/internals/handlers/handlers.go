package handlers

import (
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/handlers/venue"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/repository"
)

func NewProjectHandler() ProjectHandler {
	gormRepository := repository.NewGormRepository()
	return &projectHandlerImpl{
		venueHandler: venue.NewVenueHandler(gormRepository),
	}
}

type ProjectHandler interface {
	GetVenueHandler() *venue.VenueHandler
}

type projectHandlerImpl struct {
	venueHandler *venue.VenueHandler
}

func (ph *projectHandlerImpl) GetVenueHandler() *venue.VenueHandler {
	return ph.venueHandler
}
