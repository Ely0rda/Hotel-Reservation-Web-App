package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the applcaition config
type AppConfig struct {
	//UseCache is for specifyng from where
	//the pages will be rendered
	// from the cache
	//or CreaetemplaetCache means creating agin
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
