package controller

import (
	"encoding/json"

	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

type authController struct {
	authService contract.AuthService
}

func NewAuthController(authService contract.AuthService) contract.AuthController {
	return authController{authService}
}

func (a authController) Login(body []byte) model.LoginResponse {
	var payload model.LoginPayload

	err := common.ValidateLoginSchema(body)
	if err != nil {
		return model.LoginResponse{
			Code:    constant.BAD_REQUEST_ERROR_CODE,
			Message: err.Error(),
		}
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return model.LoginResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	return a.authService.Login(payload.Email, payload.Password)
}
