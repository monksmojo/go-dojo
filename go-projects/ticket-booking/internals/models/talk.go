package models

import "gorm.io/gorm"

type Talk struct {
	gorm.Model
	TalkName          string
	TalkType          string
	TalkDurationMins  int
	NumberOfAttendees int
	EventID           uint
	Event             Event `gorm:"foreignKey:EventID"`
	SpeakerID         uint
	Speaker           Speaker `gorm:"foreignKey:SpeakerID"`
}
