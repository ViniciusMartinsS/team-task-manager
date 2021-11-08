package contract

import "github.com/ViniciusMartinsS/manager/internal/domain"

type AuthController interface {
	Login([]byte) domain.LoginResponse
}

type AuthService interface {
	Login(email, password string) domain.LoginResponse
}
