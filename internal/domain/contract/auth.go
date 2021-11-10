package contract

import "github.com/ViniciusMartinsS/manager/internal/domain/model"

type AuthUseCases interface {
	Login([]byte) model.LoginResponse
}

type AuthService interface {
	Login(email, password string) model.LoginResponse
}
