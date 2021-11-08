package controller

import (
	"encoding/json"
	"strconv"

	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

type taskController struct {
	taskService contract.TaskService
}

func NewTaskController(taskService contract.TaskService) contract.TaskController {
	return taskController{taskService}
}

func (t taskController) List(params model.HandleTaskRequest) model.TaskResponse {
	return t.taskService.List(params.UserId)
}

func (t taskController) Create(params model.HandleTaskRequest) model.TaskResponse {
	var payload model.TaskPayload

	err := common.ValidateTaskCreateSchema(params.Body)
	if err != nil {
		return model.TaskResponse{
			Code:    constant.BAD_REQUEST_ERROR_CODE,
			Message: err.Error(),
		}
	}

	err = json.Unmarshal(params.Body, &payload)
	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	return t.taskService.Create(params.UserId, payload)
}

func (t taskController) Update(params model.HandleTaskRequest) model.TaskResponse {
	var payload model.TaskPayload

	id, err := strconv.Atoi(params.TaskId)
	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	err = json.Unmarshal(params.Body, &payload)
	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	err = common.ValidateTaskUpdateSchema(params.Body)
	if err != nil {
		return model.TaskResponse{
			Code:    constant.BAD_REQUEST_ERROR_CODE,
			Message: err.Error(),
		}
	}

	return t.taskService.Update(id, params.UserId, payload)
}

func (t taskController) Delete(params model.HandleTaskRequest) model.TaskResponse {
	id, _ := strconv.Atoi(params.TaskId)
	return t.taskService.Delete(id, params.UserId)
}
