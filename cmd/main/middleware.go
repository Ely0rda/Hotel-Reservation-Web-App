package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToconsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		//what is the next.ServeHTTP(w,r) doing
		next.ServeHTTP(w, r)
	})
}

// Nosurf adds CSRF protection to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",              //works on the entire site
		Secure:   app.InProduction, //in production it should be true
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and save the session on evry request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
