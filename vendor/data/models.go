package data

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	UserType string `json:"usertype,omitempty"`
	Email    string `json:"email"`
}

type StandardClaims struct {
	UserClaims
	jwt.StandardClaims
}
