package googleauth

import (
	"cli-todo/internal/domainErr"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type OauthService struct {
	env *Env
}

func NewOauthService(env *Env) *OauthService {
	return &OauthService{env: env}
}

// Create the redirectURl,
// if error, returns a failed redirecturl and error
func (o *OauthService) RedirectUrl() (string, error) {
	urlBuilder, err := url.Parse(o.env.AuthUrl)
	if err != nil {
		return o.env.RedirectUrlFailed, domainErr.New("failed to parse Auth url", err, domainErr.CodeInternal)
	}
	query := urlBuilder.Query()
	query.Set("response_type", "code")
	query.Set("scope", o.env.Scopes)
	query.Set("redirect_uri", o.env.RedirectUrl)
	query.Set("client_id", o.env.ClientID)
	urlBuilder.RawQuery = query.Encode()
	if urlBuilder.String() != "" {
		return o.env.RedirectUrlFailed, domainErr.New("failed to parse Auth url", err, domainErr.CodeInternal)
	}
	return urlBuilder.String(), nil
}

func (o *OauthService) ExchangeCodeForUser(url *url.URL) (*OauthUserInfo, error) {
	code := getCodeFromUrl(url)
	if code == "" {
		return nil, domainErr.New("", fmt.Errorf("faled to getCodeFromUrl"), domainErr.CodeInternal)
	}

	codeExchangeUrl := buildCodeExchangeUrl(code, o.env)
	if codeExchangeUrl == "" {
		return nil, domainErr.New("", fmt.Errorf("faled to buildCodeExchangeUrl"), domainErr.CodeInternal)
	}

	tokenInfo, err := getTokenInfo(codeExchangeUrl)
	if err != nil {
		return nil, domainErr.New("", err, domainErr.CodeInternal)
	}

	userInfo, err := getUsers(*tokenInfo)
	if err != nil {
		return nil, domainErr.New("", err, domainErr.CodeInternal)
	}
	return &userInfo, nil
}

func getCodeFromUrl(u *url.URL) string {
	urlQuery := u.Query()
	code := urlQuery.Get("code")
	if code == "" {
		return ""
	}
	return code
}

func buildCodeExchangeUrl(code string, env *Env) string {
	urlBuilder, err := url.Parse(env.TokenUrl)
	if err != nil {
		return ""
	}
	urlBuilderQuery := urlBuilder.Query()
	urlBuilderQuery.Set("code", code)
	urlBuilderQuery.Set("client_id", env.ClientID)
	urlBuilderQuery.Set("client_secret", env.ClientSecret)
	urlBuilderQuery.Set("redirect_uri", env.RedirectUrl)
	urlBuilderQuery.Set("grant_type", "authorization_code")
	urlBuilder.RawQuery = urlBuilderQuery.Encode()
	return urlBuilder.String()
}

func getTokenInfo(codeExchangeUrl string) (*OauthTokenInfo, error) {
	tokenReq, err := http.Post(codeExchangeUrl, "application/x-www-form-urlencoded", http.NoBody)
	if err != nil {
		return &OauthTokenInfo{}, err
	}
	dataBytes, err := io.ReadAll(tokenReq.Body)
	if err != nil {
		return &OauthTokenInfo{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(tokenReq.Body)

	tokenInfo := &OauthTokenInfo{}
	err = json.Unmarshal(dataBytes, tokenInfo)
	if err != nil {
		return &OauthTokenInfo{}, err
	}
	return tokenInfo, nil
}

func getUsers(t OauthTokenInfo) (OauthUserInfo, error) {
	req, err := http.NewRequest(http.MethodGet, "https://openidconnect.googleapis.com/v1/userinfo", http.NoBody)
	if err != nil {
		return OauthUserInfo{}, err
	}
	req.Header.Set("Authorization", fmt.Sprint(t.Token_type, t.Access_token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return OauthUserInfo{}, err
	}

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return OauthUserInfo{}, err
	}

	var u OauthUserInfo = OauthUserInfo{}
	err = json.Unmarshal(dataBytes, &u)
	if err != nil {
		return OauthUserInfo{}, err
	}

	//future me is going to be mad, but past me is happy
	//need to be changed if other Oauth will be added later
	u.Issuer = GoogleAuth

	return u, nil
}
