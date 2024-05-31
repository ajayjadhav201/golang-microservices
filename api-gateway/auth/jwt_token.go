package auth

import (
	"net/http"
	"strings"
	"time"

	"golang-microservices/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		common.WriteError(c, http.StatusUnauthorized, "Authorization failed")
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		common.WriteError(c, http.StatusUnauthorized, "Invalid Authorization header format")
		c.Abort()
		return
	}

	token := parts[1]
	// Validate the token (e.g., using a JWT library)
	tkn, err := ValidateToken(token)
	if err != nil {
		common.Println("auth failed error is : ", err.Error())
		common.WriteError(c, http.StatusUnauthorized, "Authorization failed")
		c.Abort()
		return
	}

	var userid = ""
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		userid = claims["userid"].(string) //user ID is stored in "user_id" claim
	}

	if userid == "" {
		common.WriteError(c, http.StatusUnauthorized, "Authorization failed, Token not valid")
		c.Abort()
		return
	}

	c.Set("userid", userid)
	c.Next()

	// ctx := context.WithValue(r.Context(), "user", claims["user"]

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
	var signingKey = []byte(common.EnvString("SIGNING_KEY", "secret_key"))

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
