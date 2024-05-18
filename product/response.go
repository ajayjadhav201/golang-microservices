package main

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewResponse(status int, data interface{}) *response {
	return &response{
		Status: status,
		Data:   data,
	}
}

func (resp *response) Write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(resp.Data)
}

// 200
func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	//
}

// 400
func StatusbadRequest(w http.ResponseWriter, r *http.Request) {
	//
}

// 404
func StatusNotFound(w http.ResponseWriter, r *http.Request) {
	//
}

// 405
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	//
}

// 409
func StatusConflict(w http.ResponseWriter, r *http.Request) {
	//
}

// 500
func StatusInternalServerError(w http.ResponseWriter, r *http.Request) {
	//
}
