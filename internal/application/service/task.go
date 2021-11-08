package service

import (
	"fmt"

	constants "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
)

type taskService struct {
	taskRepository      contract.TaskRepository
	userRepository      domain.UserRepository
	notificationService domain.NotificationService
	encryption          domain.EncryptionService
}

func NewTaskService(
	taskRepository contract.TaskRepository,
	userRepository domain.UserRepository,
	notificationService domain.NotificationService,
	encryption domain.EncryptionService,
) contract.TaskService {
	return taskService{taskRepository, userRepository, notificationService, encryption}
}

func (t taskService) List(userId int) domain.TaskResponse {
	var rows []domain.Task

	user, err := t.userRepository.FindBydId(userId)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	if constants.IsManager(user.Role.Name) {
		rows, err = t.taskRepository.FindAll()
	}

	if constants.IsTechnician(user.Role.Name) {
		rows, err = t.taskRepository.FindByUserId(userId)
	}

	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	result := make([]domain.TaskResponseContent, len(rows))

	if len(rows) == 0 {
		return domain.TaskResponse{Code: constants.SUCCESS_CODE, Result: result}
	}

	for i, r := range rows {
		result[i] = t.formatResponse(r)
	}

	return domain.TaskResponse{Code: constants.SUCCESS_CODE, Result: result}
}

func (t taskService) Create(userId int, payload domain.TaskPayload) domain.TaskResponse {
	task := domain.Task{
		Name:      payload.Name,
		Summary:   t.encryption.Encrypt(payload.Summary),
		Performed: constants.StrToDate(payload.Performed),
		UserId:    userId,
	}

	if task.Performed != nil {
		t.notificationService.Notify(task) // go
	}

	row, err := t.taskRepository.Create(task)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	result := make([]domain.TaskResponseContent, 0)
	result = append(result, t.formatResponse(row))

	return domain.TaskResponse{Code: constants.SUCCESS_CODE, Result: result}
}

func (t taskService) Update(id int, userId int, payload domain.TaskPayload) domain.TaskResponse {
	task := domain.Task{
		Name:      payload.Name,
		Summary:   payload.Summary,
		Performed: constants.StrToDate(payload.Performed),
		UserId:    userId,
	}

	if task.Summary != "" {
		task.Summary = t.encryption.Encrypt(payload.Summary)
	}

	if task.Performed != nil {
		t.notificationService.Notify(task) // go
	}

	row, err := t.taskRepository.Update(id, userId, task)

	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	result := make([]domain.TaskResponseContent, 0)
	result = append(result, t.formatResponse(row))

	return domain.TaskResponse{Code: constants.SUCCESS_CODE, Result: result}
}

func (t taskService) Delete(id int, userId int) domain.TaskResponse {
	user, err := t.userRepository.FindBydId(userId)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	if constants.IsTechnician(user.Role.Name) {
		return domain.TaskResponse{
			Code:    constants.FORBIDDEN_ERROR_CODE,
			Message: constants.FORBIDDEN_ERROR_MESSAGE,
		}
	}

	_, err = t.taskRepository.Delete(id)
	if err != nil {
		return domain.TaskResponse{
			Code:    constants.INTERNAL_SERVER_ERROR_CODE,
			Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
		}
	}

	message := fmt.Sprintf(constants.SUCCESS_DELETE_MESSAGE, id)
	return domain.TaskResponse{Code: constants.SUCCESS_CODE, Message: message}
}

func (t taskService) formatResponse(response domain.Task) domain.TaskResponseContent {
	return domain.TaskResponseContent{
		ID:        int(response.ID),
		Name:      response.Name,
		Summary:   t.encryption.Decrypt(response.Summary),
		Performed: constants.DateToStr(response.Performed),
	}
}
