package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Ely0rda/bookings/internal/config"
	"github.com/Ely0rda/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

// TestMain is the builtin function for testing
func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})
	//Change this to thrue In production
	testApp.InProduction = false
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
	session.Cookie.Secure = testApp.InProduction
	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())

}
