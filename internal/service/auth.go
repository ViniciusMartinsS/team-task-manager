package service

import (
	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/common/errors"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/ViniciusMartinsS/manager/internal/usecases/common"
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

	if err != nil && constant.DB_RECORD_NOT_FOUND == err.Error() {
		return errors.AuthNotAuthorizedErrorResponse
	}

	if err != nil {
		return errors.AuthInternalServerErrorResponse
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.AuthNotAuthorizedErrorResponse
	}

	accessToken := common.GenerateAccessToken(int(user.ID), email)
	return model.LoginResponse{Code: 0, AccessToken: accessToken}
}
