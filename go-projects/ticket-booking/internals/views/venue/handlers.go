package venue

import "net/http"

func NewVenueViewHandler() *VenueViewHandler {
	return &VenueViewHandler{}
}

type VenueViewHandler struct {
}

func (vvh *VenueViewHandler) GetVenueHomepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internals/templates/venue_homepage.html")
}

func (vvh *VenueViewHandler) GetVenueRegistrationForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internals/templates/venue_registration_form.html")
}
