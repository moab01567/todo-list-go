package googleauth

import (
	"os"
)

type Env struct {
	AuthUrl           string
	TokenUrl          string
	ClientID          string
	ClientSecret      string
	RedirectUrl       string
	Scopes            string
	RedirectUrlFailed string
}

const (
	GoogleAuthUrl        = "GOOGLE_AUTH_URL"
	GoogleClientSecret   = "GOOGLE_CLIENT_SECRET"
	GoogleClientId       = "GOOGLE_CLIENT_ID"
	GoogleRedirectUrl    = "GOOGLE_REDIRECT_URL"
	GoogleScope          = "GOOGLE_SCOPE"
	GoogleTokenUrl       = "GOOGLE_TOKEN_URL"
	GoogleRedirectFailed = "GOOGLE_REDIRECT_FAILED"
)

func NewGoogleEnv() *Env {
	env := map[string]string{
		GoogleAuthUrl:        os.Getenv(GoogleAuthUrl),
		GoogleClientSecret:   os.Getenv(GoogleClientSecret),
		GoogleClientId:       os.Getenv(GoogleClientId),
		GoogleRedirectUrl:    os.Getenv(GoogleRedirectUrl),
		GoogleScope:          os.Getenv(GoogleScope),
		GoogleTokenUrl:       os.Getenv(GoogleTokenUrl),
		GoogleRedirectFailed: os.Getenv(GoogleRedirectFailed),
	}
	for key, value := range env {
		if value == "" {
			panic("environment variable " + key + " is empty")
		}
	}

	return &Env{
		AuthUrl:           env[GoogleAuthUrl],
		TokenUrl:          env[GoogleTokenUrl],
		ClientID:          env[GoogleClientId],
		ClientSecret:      env[GoogleClientSecret],
		RedirectUrl:       env[GoogleRedirectUrl],
		Scopes:            env[GoogleScope],
		RedirectUrlFailed: env[GoogleRedirectFailed],
	}
}
