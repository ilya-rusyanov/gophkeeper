package entity

import "github.com/golang-jwt/jwt/v4"

// TokenClaims represents user auth token
type TokenClaims struct {
	jwt.RegisteredClaims
	Login string
}
