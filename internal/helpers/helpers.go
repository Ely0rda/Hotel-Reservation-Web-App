package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Ely0rda/bookings/internal/config"
)

var app *config.AppConfig

// set the app of type config.AppConfig
// using the app from the main pacakge
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	//passing the error to our InfoLog
	app.InfoLog.Println("Client error with status of ", status)
	//Writing the error to the response writer
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	//passing the error to our ErrorLog
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	//Writing the error to the response writer
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
