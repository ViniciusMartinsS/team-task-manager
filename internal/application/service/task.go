package service

import (
	"fmt"

	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

type taskService struct {
	taskRepository      contract.TaskRepository
	userRepository      contract.UserRepository
	notificationService contract.NotificationService
	encryption          contract.EncryptionService
}

func NewTaskService(
	taskRepository contract.TaskRepository,
	userRepository contract.UserRepository,
	notificationService contract.NotificationService,
	encryption contract.EncryptionService,
) contract.TaskService {
	return taskService{taskRepository, userRepository, notificationService, encryption}
}

func (t taskService) List(userId int) model.TaskResponse {
	var rows []model.Task

	user, err := t.userRepository.FindBydId(userId)
	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	if constant.IsManager(user.Role.Name) {
		rows, err = t.taskRepository.FindAll()
	}

	if constant.IsTechnician(user.Role.Name) {
		rows, err = t.taskRepository.FindByUserId(userId)
	}

	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	result := make([]model.TaskResponseContent, len(rows))

	if len(rows) == 0 {
		return model.TaskResponse{
			Code:    constant.RECORD_NOT_FOUND_ERROR_CODE,
			Message: constant.RECORD_NOT_FOUND_LIST_MESSAGE,
		}
	}

	for i, r := range rows {
		result[i] = t.formatResponse(r)
	}

	return model.TaskResponse{Code: constant.SUCCESS_CODE, Result: result}
}

func (t taskService) Create(userId int, payload model.TaskPayload) model.TaskResponse {
	task := model.Task{
		Name:      payload.Name,
		Summary:   t.encryption.Encrypt(payload.Summary),
		Performed: constant.StrToDate(payload.Performed),
		UserId:    userId,
	}

	if task.Performed != nil {
		t.notificationService.Notify(task) // go
	}

	row, err := t.taskRepository.Create(task)
	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	result := make([]model.TaskResponseContent, 0)
	result = append(result, t.formatResponse(row))

	return model.TaskResponse{Code: constant.SUCCESS_CODE, Result: result}
}

func (t taskService) Update(id int, userId int, payload model.TaskPayload) model.TaskResponse {
	task := model.Task{
		Name:      payload.Name,
		Summary:   payload.Summary,
		Performed: constant.StrToDate(payload.Performed),
		UserId:    userId,
	}

	if task.Summary != "" {
		task.Summary = t.encryption.Encrypt(payload.Summary)
	}

	if task.Performed != nil {
		t.notificationService.Notify(task) // go
	}

	row, err := t.taskRepository.Update(id, userId, task)

	if err != nil && constant.DB_RECORD_NOT_FOUND == err.Error() {
		return model.TaskResponse{
			Code:    constant.RECORD_NOT_FOUND_ERROR_CODE,
			Message: constant.RECORD_NOT_FOUND_ERROR_MESSAGE,
		}
	}

	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	result := make([]model.TaskResponseContent, 0)
	result = append(result, t.formatResponse(row))

	return model.TaskResponse{Code: constant.SUCCESS_CODE, Result: result}
}

func (t taskService) Delete(id int, userId int) model.TaskResponse {
	user, err := t.userRepository.FindBydId(userId)
	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	if constant.IsTechnician(user.Role.Name) {
		return model.TaskResponse{
			Code:    constant.FORBIDDEN_ERROR_CODE,
			Message: constant.FORBIDDEN_ERROR_MESSAGE,
		}
	}

	err = t.taskRepository.Delete(id)

	if err != nil && constant.DB_RECORD_NOT_FOUND == err.Error() {
		return model.TaskResponse{
			Code:    constant.RECORD_NOT_FOUND_ERROR_CODE,
			Message: constant.RECORD_NOT_FOUND_ERROR_MESSAGE,
		}
	}

	if err != nil {
		return model.TaskResponse{
			Code:    constant.INTERNAL_SERVER_ERROR_CODE,
			Message: constant.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	message := fmt.Sprintf(constant.SUCCESS_DELETE_MESSAGE, id)
	return model.TaskResponse{Code: constant.SUCCESS_CODE, Message: message}
}

func (t taskService) formatResponse(response model.Task) model.TaskResponseContent {
	return model.TaskResponseContent{
		ID:        int(response.ID),
		Name:      response.Name,
		Summary:   t.encryption.Decrypt(response.Summary),
		Performed: constant.DateToStr(response.Performed),
	}
}
