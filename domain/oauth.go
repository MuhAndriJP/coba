package domain

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OAUTH_URLS = []string{
	`http://localhost:9090`,
	`http://127.0.0.1:9090`,
}

var GPLUS_OAUTH_PROVIDERS map[string]*oauth2.Config

const GPLUS_CLIENTID = `134846575657-ebu1fv4euiheudn4t2v7e0l98rc1ajga.apps.googleusercontent.com`
const GPLUS_CLIENTSECRET = `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`

func init() {
	GPLUS_OAUTH_PROVIDERS = map[string]*oauth2.Config{} // yahoo tidak bisa multiple domain (harus dibuat 1-1), tidak support IP
	for _, url := range OAUTH_URLS {
		GPLUS_OAUTH_PROVIDERS[url] = &oauth2.Config{
			ClientID:     GPLUS_CLIENTID,
			ClientSecret: GPLUS_CLIENTSECRET,
			RedirectURL:  url + `/api/loginoauth`,
			Scopes: []string{
				`openid`,
				`email`,
				`profile`,
			},
			Endpoint: google.Endpoint,
		}
	}
}

func GetGPlusOAuth(host string) *oauth2.Config {
	return GPLUS_OAUTH_PROVIDERS[host]
}

const GetLink_Url = `/loginoauth`

func (d *Domain) GetLink(in *GetLink_In) (out *GetLink_Out) {

	return
}
