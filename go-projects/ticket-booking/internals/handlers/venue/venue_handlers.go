package venue

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/models"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/repository"
	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/repository/venue"
)

func NewVenueHandler(repository repository.GormRepository) *VenueHandler {
	return &VenueHandler{
		venueRepository: repository.GetVenueRepository(),
	}
}

type VenueHandler struct {
	venueRepository *venue.VenueRepository
}


func (vh *VenueHandler) CreateVenue(w http.ResponseWriter, r *http.Request) {
	log.Printf("started saving venue")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("venue_name")
	address := r.FormValue("venue_address")
	capacityStr := r.FormValue("venue_capacity")

	capacity, err := strconv.Atoi(capacityStr)
	if err != nil {
		http.Error(w, "Invalid capacity", http.StatusBadRequest)
		return
	}

	venue := models.Venue{
		VenueName:     name,
		VenueAddress:  address,
		VenueCapacity: capacity,
	}

	if err := vh.venueRepository.CreateVenue(&venue); err != nil {
		http.Error(w, "Failed to create venue", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Venue added successfully!"))
}

func (vh *VenueHandler) GetAllVenueNames(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Venue Names")
	names, err := vh.venueRepository.GetAllVenueNames()
	if err != nil {
		http.Error(w,"Failed to fetch venue names @ GetAllVenueNames",http.StatusInternalServerError)
	}

	tmpl,err:=template.ParseFiles("internals/templates/unordered_list.html")
	if err != nil {
		http.Error(w,"Failed to render unordered list template @ GetAllVenueNames",http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","text/html")
	if err

}
