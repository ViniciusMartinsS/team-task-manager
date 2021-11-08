package service

import (
	"net/http"

	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository domain.UserRepository
}

var errorNotAuthorized = "record not found"

func NewAuthService(userRepository domain.UserRepository) domain.AuthService {
	return authService{userRepository}
}

func (a authService) Login(email, password string) (domain.LoginResponse, int) {
	user, err := a.userRepository.FindByEmail(email)

	if err != nil && errorNotAuthorized == err.Error() {
		code := http.StatusUnauthorized
		return domain.LoginResponse{Message: http.StatusText(code)}, code
	}

	if err != nil {
		code := http.StatusInternalServerError
		return domain.LoginResponse{Message: http.StatusText(code)}, code
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		code := http.StatusInternalServerError
		return domain.LoginResponse{Message: http.StatusText(code)}, code
	}

	accessToken := common.GenerateAccessToken(int(user.ID), email)
	return domain.LoginResponse{Status: true, AccessToken: accessToken}, http.StatusOK
}
