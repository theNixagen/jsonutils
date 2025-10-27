// Package jsonutils
package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeJSON[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode json %w", err)
	}

	return nil
}

func DecodeJSON[T any](r *http.Request) (T, error) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode json failed: %w", err)
	}

	return data, nil
}
