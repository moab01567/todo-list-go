package httpserver

import (
	"cli-todo/internal/domainErr"
	"encoding/json"
)

func TypeToBytes(data any) ([]byte, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, domainErr.New("Server Error",
			"Could not pars type to bytes",
			err,
			domainErr.CodeInternal)
	}
	return dataBytes, nil
}
