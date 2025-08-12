package config

import "os"

const (
	GOOGLE_AUTH_URL        = "GOOGLE_AUTH_URL"
	GOOGLE_CLIENT_SECRET   = "GOOGLE_CLIENT_SECRET"
	GOOGLE_CLIENT_ID       = "GOOGLE_CLIENT_ID"
	GOOGLE_REDIRECT_URL    = "GOOGLE_REDIRECT_URL"
	GOOGLE_SCOPE           = "GOOGLE_SCOPE"
	GOOGLE_TOKEN_URL       = "GOOGLE_TOKEN_URL"
	GOOGLE_REDIRECT_FAILED = "GOOGLE_REDIRECT_FAILED"
)

type GoogleEnv struct {
	env map[string]string
}

func NewGoogleEnv() GoogleEnv {
	env := map[string]string{
		GOOGLE_AUTH_URL:        os.Getenv(string(GOOGLE_AUTH_URL)),
		GOOGLE_CLIENT_SECRET:   os.Getenv(string(GOOGLE_CLIENT_SECRET)),
		GOOGLE_CLIENT_ID:       os.Getenv(string(GOOGLE_CLIENT_ID)),
		GOOGLE_REDIRECT_URL:    os.Getenv(string(GOOGLE_REDIRECT_URL)),
		GOOGLE_SCOPE:           os.Getenv(string(GOOGLE_SCOPE)),
		GOOGLE_TOKEN_URL:       os.Getenv(string(GOOGLE_TOKEN_URL)),
		GOOGLE_REDIRECT_FAILED: os.Getenv(string(GOOGLE_REDIRECT_FAILED)),
	}
	//TODO: add ENV checker here
	return GoogleEnv{env: env}
}

func (ge GoogleEnv) GetEnv() map[string]string {
	return ge.env
}
