package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ely0rda/bookings/internal/config"
	"github.com/Ely0rda/bookings/internal/models"

	"github.com/Ely0rda/bookings/internal/render"
)

// TemplateDataholds data sent from handlers to template

// Repo is the repository variable use by handlers
var Repo *Repository

// REpository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers it sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//-----------------getting the ip and then printing it-----------
	//getting the ip address of the
	//client that send the request
	// remoteIP := r.RemoteAddr
	// //adding a key with a corresponding value
	// //to the session
	// //I don't understand the r.context()
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	//-------------------------------------------------------------------

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, mad scientist"
	//-----------------getting the ip and then printing it-----------
	// //getting the ipaddess from the session
	// remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	// //and then adding it ot stringMap
	// stringMap["remote_ip"] = remoteIP
	//--------------------------------------------------------------
	render.RenderTemplate(w, r, "about.page.html",
		&models.TemplateData{
			StringMap: stringMap,
		})

}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles requests for availability ans send json responses
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}
