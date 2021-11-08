package service

import (
	constants "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository contract.UserRepository
}

func NewAuthService(userRepository contract.UserRepository) contract.AuthService {
	return authService{userRepository}
}

func (a authService) Login(email, password string) domain.LoginResponse {
	user, err := a.userRepository.FindByEmail(email)

	if err != nil && constants.DB_ERROR_NOT_AUTHORIZED == err.Error() {
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
