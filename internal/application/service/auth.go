package service

import (
	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository contract.UserRepository
}

func NewAuthService(userRepository contract.UserRepository) contract.AuthService {
	return authService{userRepository}
}

func (a authService) Login(email, password string) model.LoginResponse {
	user, err := a.userRepository.FindByEmail(email)

	if err != nil && constant.DB_ERROR_NOT_AUTHORIZED == err.Error() {
		return model.LoginResponse{
			Code:    constant.NOT_AUTHORIZED_ERROR_CODE,
			Message: constant.NOT_AUTHORIZED_ERROR_MESSAGE,
		}
	}

	if err != nil {
		return model.LoginResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.LoginResponse{
			Code:    constant.NOT_AUTHORIZED_ERROR_CODE,
			Message: constant.NOT_AUTHORIZED_ERROR_MESSAGE,
		}
	}

	accessToken := common.GenerateAccessToken(int(user.ID), email)
	return model.LoginResponse{Code: 0, AccessToken: accessToken}
}
