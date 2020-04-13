package itswizard_siteSetup

import (
	"github.com/itslearninggermany/itswizard_jwt"
)

func (p *SiteSetup) MakeSite (admin bool) (site Site, err error) {
	if p.newAuth {
			authToken, user, err := itswizard_jwt.CreateToken(p.r,p.username,p.dbClient, p.dbWebserver)
			if err != nil {
				return site, err
			}
			p.q.Add("key",authToken)
			p.site.SessionUser = user
		} else {
			authtoken,err := itswizard_jwt.ReAuthentificate(p.r,p.dbWebserver, p.dbClient)
			if err != nil {
				return site, err
			}
			p.q.Add("key",authtoken)
		p.site.SessionUser, err = itswizard_jwt.GetUser(p.r,p.dbWebserver)
			if err != nil {
				return p.site, err
			}
		}
	p.u.RawQuery = p.q.Encode()
	p.site.URLQuerys = p.u.RawQuery
	if admin {
		p.site.Navigation = createAdminNavi(p.site.URLQuerys)
	} else {
		p.site.Navigation = createClientNavi(p.site.URLQuerys)
	}
	return p.site, err
}


func (p *SiteSetup) AddSiteName (sitename string) *SiteSetup{
	p.site.Sitename = sitename
	return p
}


func (p *SiteSetup) AddDataInSite (data interface{}) *SiteSetup{
	p.site.Special = data
	return p
}

