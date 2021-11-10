package usecases

import (
	"encoding/json"

	"github.com/ViniciusMartinsS/manager/internal/common/errors"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/ViniciusMartinsS/manager/internal/usecases/common"
)

type authUseCases struct {
	authService contract.AuthService
}

func NewAuthUseCases(authService contract.AuthService) contract.AuthUseCases {
	return authUseCases{authService}
}

func (a authUseCases) Login(body []byte) model.LoginResponse {
	var payload model.LoginPayload

	err := common.ValidateLoginSchema(body)
	if err != nil {
		return errors.AuthBadRequestErrorResponse(err.Error())
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return errors.AuthInternalServerErrorResponse
	}

	return a.authService.Login(payload.Email, payload.Password)
}
