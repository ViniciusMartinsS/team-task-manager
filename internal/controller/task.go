package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ViniciusMartinsS/manager/internal/controller/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
)

type taskController struct {
	taskService domain.TaskService
}

func NewTaskController(taskService domain.TaskService) domain.TaskController {
	return taskController{taskService}
}

func (t taskController) List(params domain.HandleTaskRequest) (domain.TaskResponse, int) {
	return t.taskService.List(params.UserId)
}

func (t taskController) Create(params domain.HandleTaskRequest) (domain.TaskResponse, int) {
	var payload domain.TaskPayload

	err := common.ValidateTaskCreateSchema(params.Body)
	if err != nil {
		code := http.StatusBadRequest
		result := domain.TaskResponse{Message: http.StatusText(code)}

		return result, code
	}

	err = json.Unmarshal(params.Body, &payload)
	if err != nil {
		code := http.StatusInternalServerError
		result := domain.TaskResponse{Message: http.StatusText(code)}

		return result, code
	}

	return t.taskService.Create(params.UserId, payload)
}

func (t taskController) Update(params domain.HandleTaskRequest) (domain.TaskResponse, int) {
	var payload domain.TaskPayload

	if params.TaskId == "" {
		code := http.StatusBadRequest
		result := domain.TaskResponse{Message: http.StatusText(code)}

		return result, code
	}

	id, _ := strconv.Atoi(params.TaskId)

	err := json.Unmarshal(params.Body, &payload)
	if err != nil {
		code := http.StatusInternalServerError
		result := domain.TaskResponse{Message: http.StatusText(code)}

		return result, code
	}

	err = common.ValidateTaskUpdateSchema(params.Body)
	if err != nil {
		code := http.StatusBadRequest
		result := domain.TaskResponse{Message: http.StatusText(code)}

		return result, code
	}

	return t.taskService.Update(id, params.UserId, payload)
}

func (t taskController) Delete(params domain.HandleTaskRequest) (domain.TaskResponse, int) {
	if params.TaskId == "" {
		code := http.StatusBadRequest
		result := domain.TaskResponse{Message: http.StatusText(code)}

		return result, code
	}

	id, _ := strconv.Atoi(params.TaskId)
	return t.taskService.Delete(id, params.UserId)
}
