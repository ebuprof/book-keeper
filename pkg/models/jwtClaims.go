package models

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	Username string `json:"username"`
	Id       string
	jwt.StandardClaims
}
