package util

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/book_keeper_go/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte(os.Getenv("JWT-SECRET"))

// this function is not in use
func GenerateToken(claims jwt.MapClaims) (string, error) {
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString([]byte("secret-key"))
}

// this function is not in use
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret_key"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid authentication token")
}

func CreateToken(username, userId string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute) // set the expiration time to 5 minutes
	claims := &models.JWTClaims{
		Username: username,
		Id:       userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenMiddleware(input http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request header
		tokenString := r.Header.Get("Authorization")

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		// I hope the token is valid
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid authorization token"))
			return
		}

		if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
			// Set username in the request context
			ctx := context.WithValue(r.Context(), "username", claims.Username)
			input(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid authorization token"))
		}
	})
}
