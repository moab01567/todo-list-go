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

func BytesToType(data []byte, v *any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return domainErr.New("Server Error",
			"Could not pars bytes to type",
			err,
			domainErr.CodeInternal)
	}
	return nil
}
