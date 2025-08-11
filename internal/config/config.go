package config

import (
	"os"
)

type GEnv string

const (
	GOOGLE_AUTH_URL        GEnv = "GOOGLE_AUTH_URL"
	GOOGLE_CLIENT_SECRET   GEnv = "GOOGLE_CLIENT_SECRET"
	GOOGLE_CLIENT_ID       GEnv = "GOOGLE_CLIENT_ID"
	GOOGLE_REDIRECT_URL    GEnv = "GOOGLE_REDIRECT_URL"
	GOOGLE_SCOPE           GEnv = "GOOGLE_SCOPE"
	GOOGLE_TOKEN_URL       GEnv = "GOOGLE_TOKEN_URL"
	GOOGLE_REDIRECT_FAILED GEnv = "GOOGLE_REDIRECT_FAILED"
)

func NewGoogleEnv() GEnv {
	return ""
}

func (ge GEnv) GetEnv() map[GEnv]string {
	return map[GEnv]string{
		GOOGLE_AUTH_URL:        os.Getenv(string(GOOGLE_AUTH_URL)),
		GOOGLE_CLIENT_SECRET:   os.Getenv(string(GOOGLE_CLIENT_SECRET)),
		GOOGLE_CLIENT_ID:       os.Getenv(string(GOOGLE_CLIENT_ID)),
		GOOGLE_REDIRECT_URL:    os.Getenv(string(GOOGLE_REDIRECT_URL)),
		GOOGLE_SCOPE:           os.Getenv(string(GOOGLE_SCOPE)),
		GOOGLE_TOKEN_URL:       os.Getenv(string(GOOGLE_TOKEN_URL)),
		GOOGLE_REDIRECT_FAILED: os.Getenv(string(GOOGLE_REDIRECT_FAILED)),
	}
}
