package auth

import (
	"fmt"
	"net/http"
)

type GoogleAuthRouter struct {
	auth_url     string
	scope        string
	redirect_url string
	client_id    string
}

func NewGoogleAuthRouter(auth_url, scope, redirect_url, client_id string) *GoogleAuthRouter {
	return &GoogleAuthRouter{
		auth_url:     auth_url,
		scope:        scope,
		redirect_url: redirect_url,
		client_id:    client_id}
}

func (g *GoogleAuthRouter) GetHandler() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /auth/google", g.redirectUser())
	handler.HandleFunc("GET /auth/google/callback", g.callback())
	return handler
}

func (g *GoogleAuthRouter) redirectUser() func(w http.ResponseWriter, r *http.Request) {
	redirectUrl := fmt.Sprintf("%s?scope=%s&&response_type=code&&redirect_uri=%s&&client_id=%s",
		g.auth_url,
		g.scope,
		g.redirect_url,
		g.client_id)

	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectUrl, http.StatusFound)
	}

}

func (g *GoogleAuthRouter) callback() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
