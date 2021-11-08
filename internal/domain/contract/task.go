package contract

import "github.com/ViniciusMartinsS/manager/internal/domain"

type TaskRepository interface {
	FindAll() ([]domain.Task, error)
	FindByUserId(id int) ([]domain.Task, error)
	Create(task domain.Task) (domain.Task, error)
	Update(id int, userId int, task domain.Task) (domain.Task, error)
	Delete(id int) (bool, error)
}

type TaskController interface {
	List(domain.HandleTaskRequest) domain.TaskResponse
	Create(domain.HandleTaskRequest) domain.TaskResponse
	Update(domain.HandleTaskRequest) domain.TaskResponse
	Delete(domain.HandleTaskRequest) domain.TaskResponse
}

type TaskService interface {
	List(userId int) domain.TaskResponse
	Create(userId int, payload domain.TaskPayload) domain.TaskResponse
	Update(id int, userId int, payload domain.TaskPayload) domain.TaskResponse
	Delete(id int, userId int) domain.TaskResponse
}
