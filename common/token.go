package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(User string) (string, error) {
	// Define the signing key (this should be kept secret)
	var signingKey = []byte(EnvString("SIGNING_KEY", "secret_key"))

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"userid":     User,
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
	var signingKey = []byte(EnvString("SIGNING_KEY", "secret_key"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, Errorf("Invalid token")
	}

	return token, nil
}
