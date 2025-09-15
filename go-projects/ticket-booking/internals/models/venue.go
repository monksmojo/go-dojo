package models

import "gorm.io/gorm"

type Venue struct {
	gorm.Model
	VenueName     string
	VenueAddress  string
	VenueCapacity int
}
