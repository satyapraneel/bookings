package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

//App config holds application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	Session       *scs.SessionManager
}
