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

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "../../templates"
	tc, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var w myWriter

	err = RenderTemplate(&w, r, "home.page.html", &models.TemplateData{})
	if err != nil {
		t.Error("Error writing template to browser")
	}
	err = RenderTemplate(&w, r, "non-existent.page.html", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that does not exist")
	}
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	//the int returned by it is the length of the slice of bytes
	length := len(b)
	return length, nil
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

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
