package itswizard_siteSetup

import (
	"github.com/itslearninggermany/itszwizard_objects"
	"html/template"
)

type Site struct {
	Navigation  template.HTML
	SessionUser itszwizard_objects.SessionUser
	Sitename    string
	Special     interface{}
	URLQuerys   string
}

