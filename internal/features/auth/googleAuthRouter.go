package auth

import (
	"cli-todo/internal/config"
	"fmt"
	"net/http"
	"os"
)

type GoogleAuthRouter struct{}

func NewGoogleAuthRouter()http.Handler{
	handler := http.NewServeMux()
	handler.HandleFunc("GET /auth/google")
	return handler
}

func (g * GoogleAuthRouter) redirectUser()func(w http.ResponseWriter, r * http.Request){
	redirectUrl := fmt.Sprintf("%s?scope=%s&&response_type=code&&redirect_uri=%s&&client_id=%s",
							os.Getenv(string(config.GOOGLE_REDIRECT_URL)),	
							os.Getenv(config.GOOGLE_SCOPE),
							os.Getenv(string(config.GOOGLE_CLIENT_ID)))
		
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w,r,os.Getenv(config.))


	}

}

