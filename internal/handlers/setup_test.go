package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/Ely0rda/bookings/internal/config"
	"github.com/Ely0rda/bookings/internal/models"
	"github.com/Ely0rda/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "../../templates"
var functions = template.FuncMap{}

// To test our handlers we need to provide them with all the necessary
// stuff they need(functions,routing,middleware,CreateTemplateCache)
// why they do it like this?
// I guess because all what we care about testing is the handlers
// and we don't want them to be affected by erros from things they don't need
// we can user CreateTemplateCache from render package
// but we can't import anything from main package
// and test wok just in the handlers directory it is like they don't have access to stuff outside it
func getRoutes() http.Handler {

	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true

	session.Cookie.SameSite = http.SameSiteLaxMode

	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")

	}

	app.TemplateCache = tc
	app.UseCache = true
	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToconsole)
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//Why we need SessionLoad be cause between
	//handlers we may want to pass somthing with
	//sessions

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)
	mux.Get("/contact", Repo.Contact)

	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/make-reservation", Repo.Reservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)

	FileServer(mux, "/static/", http.Dir("../../static"))

	return mux

}
func WriteToconsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		//what is the next.ServeHTTP(w,r) doing
		next.ServeHTTP(w, r)
	})
}

// // Nosurf adds CSRF protection to all post request
// func NoSurf(next http.Handler) http.Handler {
// 	csrfHandler := nosurf.New(next)
// 	csrfHandler.SetBaseCookie(http.Cookie{
// 		HttpOnly: true,
// 		Path:     "/",              //works on the entire site
// 		Secure:   app.InProduction, //in production it should be true
// 		SameSite: http.SameSiteLaxMode,
// 	})
// 	return csrfHandler
// }

// SessionLoad loads and save the session on evry request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {

		fs := http.StripPrefix("/static", http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
