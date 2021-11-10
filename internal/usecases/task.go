package usecases

import (
	"encoding/json"
	"strconv"

	"github.com/ViniciusMartinsS/manager/internal/common/errors"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/ViniciusMartinsS/manager/internal/usecases/common"
)

type taskUseCases struct {
	taskService contract.TaskService
}

func NewTaskUseCases(taskService contract.TaskService) contract.TaskUseCases {
	return taskUseCases{taskService}
}

func (t taskUseCases) List(params model.HandleTaskRequest) model.TaskResponse {
	return t.taskService.List(params.UserId)
}

func (t taskUseCases) Create(params model.HandleTaskRequest) model.TaskResponse {
	var payload model.TaskPayload

	err := common.ValidateTaskCreateSchema(params.Body)
	if err != nil {
		return errors.BadRequestErrorResponse(err.Error())
	}

	err = json.Unmarshal(params.Body, &payload)
	if err != nil {
		return errors.InternalServerErrorResponse
	}

	return t.taskService.Create(params.UserId, payload)
}

func (t taskUseCases) Update(params model.HandleTaskRequest) model.TaskResponse {
	var payload model.TaskPayload

	id, err := strconv.Atoi(params.TaskId)
	if err != nil {
		return errors.InternalServerErrorResponse
	}

	err = common.ValidateTaskUpdateSchema(params.Body)
	if err != nil {
		return errors.BadRequestErrorResponse(err.Error())
	}

	err = json.Unmarshal(params.Body, &payload)
	if err != nil {
		return errors.InternalServerErrorResponse
	}

	return t.taskService.Update(id, params.UserId, payload)
}

func (t taskUseCases) Delete(params model.HandleTaskRequest) model.TaskResponse {
	id, _ := strconv.Atoi(params.TaskId)
	return t.taskService.Delete(id, params.UserId)
}
