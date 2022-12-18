package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ely0rda/bookings/pkg/config"
	"github.com/Ely0rda/bookings/pkg/handlers"
	"github.com/Ely0rda/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//Change this to thrue In production
	app.InProduction = false
	// Initialize a new session manager and configure the session lifetime.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	//should the key persiste after the browser was closed by the
	// end user
	session.Cookie.Persist = true
	//how strict  yo want to be about which site this cookie
	//applies to
	session.Cookie.SameSite = http.SameSiteLaxMode
	// in production it should be true so
	//all the cookies are going to be server through
	//HTTPS
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// _ = http.ListenAndServe(portNumber, nil)
	fmt.Println(fmt.Sprintf("Starting application on port: %s ", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
