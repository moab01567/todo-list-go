package googleauth

import (
	"cli-todo/internal/domainErr"
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
		url, err := g.oauthService.RedirectUrl()
		if err != nil {
			http.Redirect(w, r, url, domainErr.GetHttpStatus(err))
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
