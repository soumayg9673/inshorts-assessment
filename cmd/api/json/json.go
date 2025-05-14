package cjson

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/soumayg9673/inshorts-assessment/cmd/api/errors"
	"github.com/soumayg9673/inshorts-assessment/internal/env"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return err
	}
	return nil
}

func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := env.GetInt("API_REQUEST_MAX_BYTES", 1_048_578)
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func WriteJSONError(w http.ResponseWriter, error errors.Error) error {
	return WriteJSON(w, error.Status, error)
}

func UnmarshalJSON(funcName string, data string, target interface{}, fieldName string) {
	if err := json.Unmarshal([]byte(data), target); err != nil {
		log.Printf("%s {%s} %v", funcName, fieldName, err)
	}
}
