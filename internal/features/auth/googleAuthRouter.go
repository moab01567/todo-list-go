package auth

import (
	"cli-todo/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GoogleEnv interface {
	GetEnv() map[config.GEnv]string
}

type GoogleAuthRouter struct {
	GoogleEnv
}

func NewGoogleAuthRouter(env GoogleEnv) *GoogleAuthRouter {
	return &GoogleAuthRouter{GoogleEnv: env}
}

func (g *GoogleAuthRouter) GetHandler() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /auth/google", g.redirectUser())
	handler.HandleFunc("GET /auth/google/callback", g.callback())
	return handler
}

func (g *GoogleAuthRouter) redirectUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlBuilder, err := url.Parse(g.GetEnv()[config.GOOGLE_AUTH_URL])
		if err != nil {
			return
		}

		query := urlBuilder.Query()
		query.Set("response_type", "code")
		query.Set("scope", g.GetEnv()[config.GOOGLE_SCOPE])
		query.Set("redirect_uri", g.GetEnv()[config.GOOGLE_REDIRECT_URL])
		query.Set("client_id", g.GetEnv()[config.GOOGLE_CLIENT_ID])
		urlBuilder.RawQuery = query.Encode()

		http.Redirect(w, r, urlBuilder.String(), http.StatusFound)

	}
}

func (g *GoogleAuthRouter) callback() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		code := getCodeFromUrl(r.URL)
		if code == "" {
			return
		}

		codeExchangeUrl := g.buildCodeExchangeUrl(code)
		if codeExchangeUrl == "" {
			return
		}

		t, err := getTokenInfo(codeExchangeUrl)
		if err != nil {
			return
		}

		u, err := getUsers(t)
		if err != nil {
			return
		}
		fmt.Println(u)

	}
}

// -------helpers---------
type tokenInfo struct {
	Token_type   string
	Access_token string
	Scope        string
	Id_token     string
}

func getCodeFromUrl(u *url.URL) string {
	urlQuery := u.Query()
	code := urlQuery.Get("code")
	if code == "" {
		return ""
	}
	return code
}

func (g *GoogleAuthRouter) buildCodeExchangeUrl(code string) string {
	urlBuilder, err := url.Parse(g.token_url)
	if err != nil {
		return ""
	}
	urlBuilderQuery := urlBuilder.Query()
	urlBuilderQuery.Set("code", code)
	urlBuilderQuery.Set("client_id", g.client_id)
	urlBuilderQuery.Set("client_secret", g.client_secret)
	urlBuilderQuery.Set("redirect_uri", g.redirect_url)
	urlBuilderQuery.Set("grant_type", "authorization_code")
	urlBuilder.RawQuery = urlBuilderQuery.Encode()
	return urlBuilder.String()
}

func getTokenInfo(codeExchangeUrl string) (tokenInfo, error) {
	tokenReq, err := http.Post(codeExchangeUrl, "application/x-www-form-urlencoded", http.NoBody)
	if err != nil {
		return tokenInfo{}, err
	}
	dataBytes, err := io.ReadAll(tokenReq.Body)
	if err != nil {
		return tokenInfo{}, err
	}

	var t tokenInfo = tokenInfo{}
	err = json.Unmarshal(dataBytes, &t)
	if err != nil {
		return tokenInfo{}, err
	}
	return t, nil
}

type userInfo struct {
	Sub     string
	Name    string
	Email   string
	Picture string
}

func getUsers(t tokenInfo) (userInfo, error) {
	req, err := http.NewRequest(http.MethodGet, "https://openidconnect.googleapis.com/v1/userinfo", http.NoBody)
	if err != nil {
		return userInfo{}, err
	}
	req.Header.Set("Authorization", fmt.Sprint(t.Token_type, t.Access_token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return userInfo{}, err
	}

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return userInfo{}, err
	}

	var u userInfo = userInfo{}
	err = json.Unmarshal(dataBytes, &u)
	if err != nil {
		return userInfo{}, err
	}

	return u, nil
}
