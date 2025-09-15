package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventName         string
	EventStartDate    time.Time
	EventEndDate      time.Time
	EventPrice        float64
	NumberOfAttendees int
	VenueID           uint
	Venue             Venue  `gorm:"foreignKey:VenueID"`
	Talks             []Talk `gorm:"foreignKey:EventID;constraint:OnDelete:SET NULL"`
}
