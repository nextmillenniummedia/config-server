package utils

import (
	"errors"
	"io"
	"net/http"
)

func ReadPayload(r *http.Request) ([]byte, error) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if len(payload) == 0 {
		return nil, errors.New("empty request payload")
	}
	return payload, nil
}
