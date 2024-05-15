package common

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	UnmarshalFailed   = "Failed to parse request body"
	InternalServerErr = "Internal server error"
)

var validate *validator.Validate = validator.New()

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
	if err == nil {
		return nil
	}
	//create custom error messages here
	newerr := ""
	if strings.Contains(err.Error(), "Key") {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			fieldTag := err.Tag()
			Println("ajaj fields are name: ", fieldName, "tag : ", fieldTag)
			switch fieldTag {
			case "email":
				newerr = SPrintf("%s, %s", newerr, "email is not valid")

			case "required":
				newerr = SPrintf("%s, %s is required", newerr, fieldName)
			}
		}
		return errors.New(newerr)
	}
	return err

	// if r.ContentLength <= 0 {
	// 	return errors.New("bad request")
	// }
	// b, e := io.ReadAll(r.Body)
	// if e != nil {
	// 	return errors.New("failed to read request body")
	// }
	// return json.Unmarshal(b, pointer)
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
