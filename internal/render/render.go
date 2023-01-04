package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Ely0rda/bookings/internal/config"
	"github.com/Ely0rda/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates set the config for the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	//td.Flash wil be a message shown when a everything went well"success"
	td.Flash = app.Session.PopString(r.Context(), "flash")
	//td.Error will be a message shown when there is an error
	td.Error = app.Session.PopString(r.Context(), "error")
	//td.Wrning will be a message shown as a warning
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) { //, td *models.TemplateData
// 	var tc map[string]*template.Template

// 	if app.UseCache {
// 		tc = app.TemplateCache
// 		// fmt.Println("did not use the cache", app.UseCache)

// 	} else {
// 		// 	fmt.Println("the createTemplateCache was calleld", app.UseCache)

// 		tc, _ = CreateTemplateCache()

// 	}

// 	t, ok := tc[tmpl]
// 	//apparently if the map can't find the given key
// 	//it will return a bool false which we have used here
// 	//as ok
// 	if !ok {
// 		log.Fatal("render pakcage: could not get the templates from the template cache ", len(app.TemplateCache))
// 	}
// 	//according to mr dumb?
// 	//in this version of code before
// 	//we were getting the templates directly from the disc
// 	//as html files and then we do all the neessary operations
// 	// parsedTemplate, _ := template.ParseFiles("../../templates/" + tmpl)
// 	// err := parsedTemplate.Execute(w, nil)
// 	// if err != nil {
// 	// 	fmt.Println("error parsing template : ", err)
// 	// 	try to delete the return
// 	// 	return
// 	// }
// 	//but now the template we did not get the template from
// 	//the disk so we need to handle it diffrently now
// 	//we need to createa a byte buffer a variable
// 	// type for holding bytes
// 	//and we will put the parsed template t in it after being
// 	//executed
// 	//me playing with the code---------------------
// 	buf := new(bytes.Buffer)
// 	td = AddDefaultData(td)
// 	_ = t.Execute(buf, td)
// 	_, err := buf.WriteTo(w)
// 	if err != nil {
// 		fmt.Println("Error writing the template to the browser")
// 	}
// 	// -----------------------------------------------
// 	//the problme with the code befor is that it can
// 	//be siplfy and be written just like this
// 	//because execute takes a *template.Template
// 	//and than pare it and write it to w
// 	//func (t *Template) Execute(wr io.Writer, data any)
// 	//and t is already a *template.Template type so
// 	//no problem
// 	// err = t.Execute(w, nil)
// 	// if err != nil {
// 	// 	fmt.Println("error executing the template")
// 	// }
// }

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("RenderTemplate Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

// CreateTemplateCache is for creating template cache as a map
// wich means it creates all the template so tey can be ready
// to use when we need them.
// func CreateTemplateCache() (map[string]*template.Template, error) {
// 	//myCach hold all of the templates
// 	//and it creates them at the start of the applciation
// 	myCache := map[string]*template.Template{}
// 	//Glob returns the names of all files matching pattern or nil
// 	//getting paths to all the pages with a name ending with .page.html
// 	// from the ../../templates path
// 	pages, err := filepath.Glob("../../templates/*.page.html")
// 	if err != nil {
// 		return myCache, err
// 	}
// 	//The whole point from the loop below is to fill myCach
// 	//So we are going to go through all the pages(templates) and
// 	//and parse those templates while adding their
// 	//layouts and then, insert them in the myCach[nameOfTemplate,templateParsed]
// 	for _, page := range pages {
// 		//Base return the last element of the path
// 		//now here we are trying to determine just the name of pages
// 		//from the paths and storing them in name
// 		fmt.Println("Page is currently", pages)
// 		name := filepath.Base(page)

// 		//New allocates a new undefined template  with
// 		// the given name
// 		//Funcs adds the elements of the argument map to the template's function map.
// 		//my basic understandingis that every template in golang has
// 		//function map and we are going to add functions to this
// 		//function map
// 		//.ParseFiles parses the named files and associates
// 		// the resulting templates, with the template New
// 		//is going to return
// 		//Now ts should cotain a set of templates
// 		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
// 		if err != nil {
// 			return myCache, err
// 		}

// 		//look in the ../../templates directory
// 		//for any files that ends with layout.html
// 		matches, err := filepath.Glob("../../templates/*.layout.html")
// 		if err != nil {
// 			return myCache, err
// 		}
// 		//checking if matches contain something or not
// 		if len(matches) > 0 {
// 			// ParseGlob parses the template definitions in the files identified by the
// 			// pattern and associates the resulting templates with t.
// 			//func (t *Template) ParseGlob(pattern string) (*Template, error)
// 			//look for any matching layout in the given path and then
// 			//parse it to ts
// 			ts, err = ts.ParseGlob("../../templates/*.layout.html")
// 			if err != nil {
// 				return myCache, err
// 			}
// 		}
// 		myCache[name] = ts

// 	}

// 	return myCache, nil
// }

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
