package repository

import (
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/db"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/repository/venue"
)

func NewGormRepository() GormRepository {
	return &gormRepositoryImpl{
		venueRepository: venue.NewVenueRepository(db.DB),
	}
}

type GormRepository interface {
	GetVenueRepository() *venue.VenueRepository
}

type gormRepositoryImpl struct {
	venueRepository *venue.VenueRepository
}

func (gr *gormRepositoryImpl) GetVenueRepository() *venue.VenueRepository {
	return gr.venueRepository
}
