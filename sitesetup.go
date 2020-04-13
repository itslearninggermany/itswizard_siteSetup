package itswizard_siteSetup

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"net/url"
)

type SiteSetup struct {
	u           *url.URL
	q           *url.Values
	jwt         string
	site        Site
	dbWebserver *gorm.DB
	dbClient    *gorm.DB
	r           *http.Request
	newAuth     bool
	username    string
}

func InitialSite(scheme, host string, newAuth bool, username string) *SiteSetup {
	s := new(SiteSetup)
	s.u.Scheme = scheme
	s.u.Host = host
	s.newAuth = newAuth
	return s
}
