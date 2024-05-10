package main

import (
	"net/http"

	"github.com/ajayjadhav201/common"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var token string

	// Extract JWT token from the request based on the client type
	if isWebAppRequest(r) {
		// Extract JWT token from HTTP-only cookie for web applications
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			http.Error(w, "JWT token not found", http.StatusUnauthorized)
			return
		}
		token = cookie.Value
	} else if isAndroidAppRequest(r) {
		// Extract JWT token from Authorization header for Android applications
		token = extractTokenFromAuthorizationHeader(r)
	} else {
		http.Error(w, "Unsupported client", http.StatusUnauthorized)
		return
	}

	common.Print(token)

	// Use the JWT token as needed...
}

// Check if the request is from a web application
func isWebAppRequest(r *http.Request) bool {
	// Implement logic to identify requests from web applications
	return false
}

// Check if the request is from an Android application
func isAndroidAppRequest(r *http.Request) bool {
	// Implement logic to identify requests from Android applications
	return false
}

// Extract JWT token from Authorization header
func extractTokenFromAuthorizationHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	common.Print(authHeader)
	return authHeader
	// Extract token from the Authorization header (e.g., "Bearer <JWT_token>")
	// Parse the token and return it
}
