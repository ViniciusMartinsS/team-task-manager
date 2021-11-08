package model

import "github.com/golang-jwt/jwt"

type LoginPayload struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type LoginResponse struct {
	Code        int    `json:"status"`
	AccessToken string `json:"accessToken,omitempty"`
	Message     string `json:"message,omitempty"`
}

type Claims struct {
	UserId int    `json:"userId"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
