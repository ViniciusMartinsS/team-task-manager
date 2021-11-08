package controller

import (
	"encoding/json"

	constants "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
)

type authController struct {
	authService contract.AuthService
}

func NewAuthController(authService contract.AuthService) contract.AuthController {
	return authController{authService}
}

func (a authController) Login(body []byte) domain.LoginResponse {
	var payload domain.LoginPayload

	err := common.ValidateLoginSchema(body)
	if err != nil {
		return domain.LoginResponse{
			Code:    constants.BAD_REQUEST_ERROR_CODE,
			Message: err.Error(),
		}
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return domain.LoginResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	return a.authService.Login(payload.Email, payload.Password)
}
