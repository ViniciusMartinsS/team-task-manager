package service

import (
	constants "github.com/ViniciusMartinsS/manager/internal/common"
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

func (a authService) Login(email, password string) domain.LoginResponse {
	user, err := a.userRepository.FindByEmail(email)

	if err != nil && errorNotAuthorized == err.Error() {
		return domain.LoginResponse{
			Code:    constants.NOT_AUTHORIZED_ERROR_CODE,
			Message: constants.NOT_AUTHORIZED_ERROR_MESSAGE,
		}
	}

	if err != nil {
		return domain.LoginResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.LoginResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	accessToken := common.GenerateAccessToken(int(user.ID), email)
	return domain.LoginResponse{Code: 0, AccessToken: accessToken}
}
