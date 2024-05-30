package auth

import (
	"net/http"
	"strings"
	"time"

	"golang-microservices/common"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	//
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		// Validate the token (e.g., using a JWT library)
		_, err := ValidateToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// ctx := context.WithValue(r.Context(), "user", claims["user"])
		next.ServeHTTP(w, r)
	})
}

func CreateToken(UserId string) (string, error) {
	// Define the signing key (this should be kept secret)
	var signingKey = []byte(common.EnvString("SIGNING_KEY", "secret_key"))

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user":       UserId,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	var signingKey = []byte("your-256-bit-secret")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, common.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, common.Errorf("invalid token")
	}

	return token, nil
}