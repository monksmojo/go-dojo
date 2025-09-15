package models

import "gorm.io/gorm"

type Speaker struct {
	gorm.Model
	SpeakerName string
	SpeakerAge  string
	Talk        []Talk `gorm:"foreignKey:SpeakerID;constraint:OnDelete:CASCADE"`
}
