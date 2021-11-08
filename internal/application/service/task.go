package service

import (
	"net/http"

	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/helper"
)

type taskService struct {
	taskRepository      domain.TaskRepository
	userRepository      domain.UserRepository
	notificationService domain.NotificationService
	encryption          domain.Encryption
}

func NewTaskService(
	taskRepository domain.TaskRepository,
	userRepository domain.UserRepository,
	notificationService domain.NotificationService,
	encryption domain.Encryption,
) domain.TaskService {
	return taskService{taskRepository, userRepository, notificationService, encryption}
}

func (t taskService) List(userId int) (domain.TaskResponse, int) {
	var rows []domain.Task

	user, err := t.userRepository.FindBydId(userId)
	if err != nil {
		code := http.StatusInternalServerError
		return domain.TaskResponse{Message: http.StatusText(code)}, code
	}

	if helper.IsManager(user.Role.Name) {
		rows, err = t.taskRepository.FindAll()
	}

	if helper.IsTechnician(user.Role.Name) {
		rows, err = t.taskRepository.FindByUserId(userId)
	}

	if err != nil {
		code := http.StatusInternalServerError
		return domain.TaskResponse{Message: http.StatusText(code)}, code
	}

	code := http.StatusOK
	result := make([]domain.TaskResponseContent, len(rows))

	if len(rows) == 0 {
		return domain.TaskResponse{Status: true, Result: result}, code
	}

	for i, r := range rows {
		result[i] = t.formatResponse(r)
	}

	return domain.TaskResponse{Status: true, Result: result}, code
}

func (t taskService) Create(userId int, payload domain.TaskPayload) (domain.TaskResponse, int) {
	task := domain.Task{
		Name:      payload.Name,
		Summary:   t.encryption.Encrypt(payload.Summary),
		Performed: helper.StrToDate(payload.Performed),
		UserId:    userId,
	}

	if task.Performed != nil {
		t.notificationService.Notify(task) // go
	}

	row, err := t.taskRepository.Create(task)
	if err != nil {
		code := http.StatusInternalServerError
		return domain.TaskResponse{Message: http.StatusText(code)}, code
	}

	result := make([]domain.TaskResponseContent, 0)
	result = append(result, t.formatResponse(row))

	return domain.TaskResponse{Status: true, Result: result}, http.StatusCreated
}

func (t taskService) Update(id int, userId int, payload domain.TaskPayload) (domain.TaskResponse, int) {
	task := domain.Task{
		Name:      payload.Name,
		Summary:   payload.Summary,
		Performed: helper.StrToDate(payload.Performed),
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
		code := http.StatusInternalServerError
		return domain.TaskResponse{Message: http.StatusText(code)}, code
	}

	result := make([]domain.TaskResponseContent, 0)
	result = append(result, t.formatResponse(row))

	return domain.TaskResponse{Status: true, Result: result}, http.StatusCreated
}

func (t taskService) Delete(id int, userId int) (domain.TaskResponse, int) {
	user, err := t.userRepository.FindBydId(userId)
	if err != nil {
		code := http.StatusInternalServerError
		return domain.TaskResponse{Message: http.StatusText(code)}, code
	}

	if helper.IsTechnician(user.Role.Name) {
		code := http.StatusForbidden
		return domain.TaskResponse{Message: http.StatusText(code)}, code
	}

	_, err = t.taskRepository.Delete(id)
	if err != nil {
		code := http.StatusInternalServerError
		return domain.TaskResponse{Message: http.StatusText(code)}, code
	}

	code := http.StatusOK
	return domain.TaskResponse{Message: "Register Deleted."}, code
}

func (t taskService) formatResponse(response domain.Task) domain.TaskResponseContent {
	return domain.TaskResponseContent{
		ID:        int(response.ID),
		Name:      response.Name,
		Summary:   t.encryption.Decrypt(response.Summary),
		Performed: helper.DateToStr(response.Performed),
	}
}
