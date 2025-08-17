package googleauth

import (
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

func (o OauthService) ExchangeCodeForUser(url *url.URL) (*OauthUserInfo, error) {
	code := getCodeFromUrl(url)
	if code == "" {
		fmt.Println("her1")
		//Todo handle error
		return nil, nil
	}

	codeExchangeUrl := buildCodeExchangeUrl(code, o.env)
	if codeExchangeUrl == "" {
		//Todo handle error
		fmt.Println("her2")
		return nil, nil
	}

	tokenInfo, err := getTokenInfo(codeExchangeUrl)
	if err != nil {
		//Todo handle error
		fmt.Println("her3")
		fmt.Println(err)
		return nil, nil
	}

	userInfo, err := getUsers(*tokenInfo)
	if err != nil {
		//Todo handle error
		fmt.Println("her4")
		return nil, nil
	}
	fmt.Println(userInfo)
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

	fmt.Println(string(dataBytes))
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

	return u, nil
}
