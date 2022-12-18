package handlers

import (
	"net/http"

	"github.com/Ely0rda/bookings/pkg/config"
	"github.com/Ely0rda/bookings/pkg/models"

	"github.com/Ely0rda/bookings/pkg/render"
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

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
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
	render.RenderTemplate(w, "about.page.html",
		&models.TemplateData{
			StringMap: stringMap,
		})

}
