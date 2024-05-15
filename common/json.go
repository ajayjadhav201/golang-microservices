package common

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	UnmarshalFailed   = "Failed to parse request body"
	InternalServerErr = "Internal server error"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ReadJSON(r *http.Request, data any) error {
	// return json.NewDecoder(r.Body).Decode(data)
	if r.ContentLength <= 0 {
		return errors.New("bad request")
	}
	b, e := io.ReadAll(r.Body)
	if e != nil {
		return errors.New("failed to read request body")
	}
	return json.Unmarshal(b, &data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}

func WriteInternalServerError(w http.ResponseWriter) {
	WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error."})
}

func WriteServerNotAvailableError(w http.ResponseWriter) {
	WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Server not available."})
}
