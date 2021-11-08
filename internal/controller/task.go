package controller

import (
	"encoding/json"
	"strconv"

	constants "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
)

type taskController struct {
	taskService contract.TaskService
}

func NewTaskController(taskService contract.TaskService) contract.TaskController {
	return taskController{taskService}
}

func (t taskController) List(params domain.HandleTaskRequest) domain.TaskResponse {
	return t.taskService.List(params.UserId)
}

func (t taskController) Create(params domain.HandleTaskRequest) domain.TaskResponse {
	var payload domain.TaskPayload

	err := common.ValidateTaskCreateSchema(params.Body)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.BAD_REQUEST_ERROR_CODE,
			Message: err.Error(),
		}
	}

	err = json.Unmarshal(params.Body, &payload)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	return t.taskService.Create(params.UserId, payload)
}

func (t taskController) Update(params domain.HandleTaskRequest) domain.TaskResponse {
	var payload domain.TaskPayload

	id, err := strconv.Atoi(params.TaskId)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	err = json.Unmarshal(params.Body, &payload)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	err = common.ValidateTaskUpdateSchema(params.Body)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.BAD_REQUEST_ERROR_CODE,
			Message: err.Error(),
		}
	}

	return t.taskService.Update(id, params.UserId, payload)
}

func (t taskController) Delete(params domain.HandleTaskRequest) domain.TaskResponse {
	id, _ := strconv.Atoi(params.TaskId)
	return t.taskService.Delete(id, params.UserId)
}
