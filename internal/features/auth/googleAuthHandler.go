package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GoogleAuthRouter struct {
	env map[GoogleEnv]string
}

func NewGoogleAuthRouter() *GoogleAuthRouter {
	return &GoogleAuthRouter{env: GetGoogleEnvMap()}
}

func (g *GoogleAuthRouter) GetHandler() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /auth/google", g.redirectUser())
	handler.HandleFunc("GET /auth/google/callback", g.callback())
	handler.HandleFunc("GET /auth/success", g.callback())
	return handler
}

func (g *GoogleAuthRouter) redirectUser() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		urlBuilder, err := url.Parse(g.env[GoogleAuthUrl])
		if err != nil {
			return
		}
		query := urlBuilder.Query()
		query.Set("response_type", "code")
		query.Set("scope", g.env[GoogleScope])
		query.Set("redirect_uri", g.env[GoogleRedirectUrl])
		query.Set("client_id", g.env[GoogleClientId])
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
		http.SetCookie(w, &http.Cookie{Name: "code",
			Value:    "test",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
			MaxAge:   60})
		w.Write([]byte("hello world " + u.Name))
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
	urlBuilder, err := url.Parse(g.env[GoogleAuthUrl])
	if err != nil {
		return ""
	}
	urlBuilderQuery := urlBuilder.Query()
	urlBuilderQuery.Set("code", code)
	urlBuilderQuery.Set("client_id", g.env[GoogleClientId])
	urlBuilderQuery.Set("client_secret", g.env[GoogleClientSecret])
	urlBuilderQuery.Set("redirect_uri", g.env[GoogleRedirectUrl])
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
