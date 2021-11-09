package errors

import (
	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

var InternalServerErrorResponse = model.TaskResponse{
	Code:    constant.INTERNAL_SERVER_ERROR_CODE,
	Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
}

var ListRecordNotFoundErrorResponse = model.TaskResponse{
	Code:    constant.RECORD_NOT_FOUND_ERROR_CODE,
	Message: constant.RECORD_NOT_FOUND_LIST_MESSAGE,
}

var RecordNotFoundErrorResponse = model.TaskResponse{
	Code:    constant.RECORD_NOT_FOUND_ERROR_CODE,
	Message: constant.RECORD_NOT_FOUND_ERROR_MESSAGE,
}

var BadRequestErrorResponse = func(err string) model.TaskResponse {
	return model.TaskResponse{
		Code:    constant.BAD_REQUEST_ERROR_CODE,
		Message: err,
	}
}
