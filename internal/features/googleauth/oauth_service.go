package googleauth

import "net/url"

type OauthService struct {
	env *Env
}

func NewOauthService(env *Env) *OauthService {
	return &OauthService{env: env}
}

func (o *OauthService) RedirectUrl() string {
	urlBuilder, err := url.Parse(o.env.AuthUrl)
	if err != nil {
		return ""
	}
	query := urlBuilder.Query()
	query.Set("response_type", "code")
	query.Set("scope", o.env.Scopes)
	query.Set("redirect_uri", o.env.RedirectUrl)
	query.Set("client_id", o.env.ClientID)
	urlBuilder.RawQuery = query.Encode()

	return urlBuilder.String()
}
