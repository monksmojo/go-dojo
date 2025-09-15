package venue

import (
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/models"
	"gorm.io/gorm"
)

type VenueRepository struct {
	db *gorm.DB
}

func (vr *VenueRepository) CreateVenue(venue *models.Venue) error {
	err := vr.db.Create(venue).Error
	if err != nil {
		return err
	}

	return nil
}

func (vr *VenueRepository) GetAllVenueNames() ([]string, error) {
	var names []string
	err := vr.db.Model(&models.Venue{}).Pluck("venue_name", &names).Error
	return names, err
}

func NewVenueRepository(db *gorm.DB) *VenueRepository {
	return &VenueRepository{
		db: db,
	}
}
