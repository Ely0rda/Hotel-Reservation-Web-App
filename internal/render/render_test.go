package render

import (
	"net/http"
	"testing"

	"github.com/Ely0rda/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	//func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData
	//AddDefaultData takes a templateData and an http.Request
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	//putting data into the session (the session is in the request)
	session.Put(r.Context(), "flash", "123")
	// func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData
	//AddDefaultData retrieve some data from the sessions and returns TemplateData
	//and if it did not find any it is going to panic
	//and the test is going to fail
	//Passing the request and templateData to it, and getting the templateData
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session ")
	}

}

func getSession() (*http.Request, error) {
	//Creating a request
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	//getting the request context
	ctx := r.Context()
	//I don't know what the following code does
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	//With context returns a copy of r with the given context.
	r = r.WithContext(ctx)
	return r, nil
}
