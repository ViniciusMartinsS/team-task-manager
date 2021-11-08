package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
)

type authController struct {
	authService domain.AuthService
}

func NewAuthController(authService domain.AuthService) domain.AuthController {
	return authController{authService}
}

func (a authController) Login(body []byte) (domain.LoginResponse, int) {
	var payload domain.LoginPayload

	err := common.ValidateLoginSchema(body)
	if err != nil {
		code := http.StatusBadRequest
		result := domain.LoginResponse{Message: http.StatusText(code)}

		return result, code
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		code := http.StatusInternalServerError
		result := domain.LoginResponse{Message: http.StatusText(code)}

		return result, code
	}

	return a.authService.Login(payload.Email, payload.Password)
}
