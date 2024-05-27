package common

import (
	"encoding/json"
	"net/http"
)

const (
	UnmarshalFailed   = "Failed to parse request body"
	InternalServerErr = "Internal server error"
)

type Response struct {
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ReadJSON(r *http.Request, pointer any) error {
	err := json.NewDecoder(r.Body).Decode(pointer)
	if err != nil {
		return err
	}
	err = validate.Struct(pointer)
	if err != nil {
		return err
	}
	return nil

	// if r.ContentLength <= 0 {
	// 	return errors.New("bad request")
	// }
	// b, e := io.ReadAll(r.Body)
	// if e != nil {
	// 	return errors.New("failed to read request body")
	// }
	// return json.Unmarshal(b, pointer)
}
