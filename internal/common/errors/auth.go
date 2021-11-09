package errors

import (
	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

var AuthInternalServerErrorResponse = model.LoginResponse{
	Code:    constant.INTERNAL_SERVER_ERROR_CODE,
	Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
}

var AuthNotAuthorizedErrorResponse = model.LoginResponse{
	Code:    constant.NOT_AUTHORIZED_ERROR_CODE,
	Message: constant.NOT_AUTHORIZED_ERROR_MESSAGE,
}

var AuthBadRequestErrorResponse = func(err string) model.LoginResponse {
	return model.LoginResponse{
		Code:    constant.BAD_REQUEST_ERROR_CODE,
		Message: err,
	}
}
