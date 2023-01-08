package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ely0rda/bookings/internal/config"
	"github.com/Ely0rda/bookings/internal/forms"
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		//The template contains value from the elements below
		//this why we should pass them,even if those elements
		//are empty, because if we didn't the template will not
		//be parsed

		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm  parses the form
	//and allow us to use functions r.Form.Get, so we can check the data that
	//was sent
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	reservation := models.Reservation{
		//r.Form contains the parsed from data
		//Only available after r.ParseForm is
		//called
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}
	//r.PostForm contains the parsed form data
	form := forms.New(r.PostForm)
	//passing all the fields we want to validate to
	//form.Required
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	_ = form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	//we want passe the reservation var to another template
	//and to achieve that we will be using sessions
	//passing the reservation through sessions
	m.App.Session.Put(r.Context(), "reservation", reservation)
	//now becaue everything went well we will
	//redirect the user to the resrvation summar page
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	//getting the resrvatiom from the session
	// and checking its type **.(models.Reservation)**

	//if the reservation variable was not in the session the template will not
	//rendered
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		//If the resrvation var is not in the session this means that
		//probably the user pass to this page before make-resrvation
		//this why we will create an error to be shown in the next
		//page we will redirect the user to
		//which is going to be the home page
		m.App.Session.Put(r.Context(), "error", "Can't get resrvation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "reservation")
	//making a data variable
	data := make(map[string]interface{})
	//adding the resrvation to our data variable
	data["reservation"] = reservation
	//Calling the renderTemplate and passing data to the template
	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
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
