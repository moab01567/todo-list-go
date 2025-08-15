package auth

import "os"

type GoogleEnv string

const (
	GoogleAuthUrl        GoogleEnv = "GOOGLE_AUTH_URL"
	GoogleClientSecret   GoogleEnv = "GOOGLE_CLIENT_SECRET"
	GoogleClientId       GoogleEnv = "GOOGLE_CLIENT_ID"
	GoogleRedirectUrl    GoogleEnv = "GOOGLE_REDIRECT_URL"
	GoogleScope          GoogleEnv = "GOOGLE_SCOPE"
	GoogleTokenUrl       GoogleEnv = "GOOGLE_TOKEN_URL"
	GoogleRedirectFailed GoogleEnv = "GOOGLE_REDIRECT_FAILED"
)

func GetGoogleEnvMap() map[GoogleEnv]string {
	env := map[GoogleEnv]string{
		GoogleAuthUrl:        os.Getenv(string(GoogleAuthUrl)),
		GoogleClientSecret:   os.Getenv(string(GoogleClientSecret)),
		GoogleClientId:       os.Getenv(string(GoogleClientId)),
		GoogleRedirectUrl:    os.Getenv(string(GoogleRedirectUrl)),
		GoogleScope:          os.Getenv(string(GoogleScope)),
		GoogleTokenUrl:       os.Getenv(string(GoogleTokenUrl)),
		GoogleRedirectFailed: os.Getenv(string(GoogleRedirectFailed)),
	}
	for key, value := range env {
		if value == "" {
			panic("environment variable " + key + " is empty")
		}
	}
	return env
}
