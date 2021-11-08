package helper

import (
	"time"

	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret_key")

var key = func(token *jwt.Token) (interface{}, error) {
	return jwtKey, nil
}

func IsAccessTokenValid(accessToken string) (bool, *domain.Claims) {
	claims := &domain.Claims{}
	token, _ := jwt.ParseWithClaims(accessToken, claims, key)
	return token.Valid, claims
}

func GenerateAccessToken(id int, email string) string {
	expirationTime := time.Now().Add(150 * time.Minute)

	claims := &domain.Claims{
		Email:  email,
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, _ := token.SignedString(jwtKey)

	return accessToken
}
