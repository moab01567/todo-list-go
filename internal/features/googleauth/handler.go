package googleauth

import (
	"net/http"
)

type GoogleHandler struct {
	oauthService *OauthService
	authService  AuthService
}

func NewGoogleHandler(oauthService *OauthService, authService AuthService) *GoogleHandler {
	return &GoogleHandler{oauthService: oauthService, authService: authService}
}

func (g *GoogleHandler) Routes() http.Handler {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("GET /auth/google", g.redirectUser())
	serveMux.HandleFunc("GET /auth/google/callback", g.callback())
	return serveMux
}

func (g *GoogleHandler) redirectUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := g.oauthService.RedirectUrl()
		if url == "" {
			//Todo Handle Error
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func (g *GoogleHandler) callback() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo, err := g.oauthService.ExchangeCodeForUser(r.URL)
		if err != nil {
			//Todo handle error
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "code",
			Value:    "test",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
			MaxAge:   60})
		//Todo Handle error
		w.Write([]byte("hello world " + userInfo.Name))

	}
}
