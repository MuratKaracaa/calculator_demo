package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func DecodeAndValidate[T Validator](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		var zero T
		return zero, fmt.Errorf("invalid body: %w", err)
	}
	if err := v.Validate(); err != nil {
		var zero T
		return zero, fmt.Errorf("validation failed: %w", err)
	}
	return v, nil
}
