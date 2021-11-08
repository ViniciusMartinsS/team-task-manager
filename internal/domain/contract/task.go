package contract

import "github.com/ViniciusMartinsS/manager/internal/domain/model"

type TaskRepository interface {
	FindAll() ([]model.Task, error)
	FindByUserId(id int) ([]model.Task, error)
	Create(task model.Task) (model.Task, error)
	Update(id int, userId int, task model.Task) (model.Task, error)
	Delete(id int) (bool, error)
}

type TaskController interface {
	List(model.HandleTaskRequest) model.TaskResponse
	Create(model.HandleTaskRequest) model.TaskResponse
	Update(model.HandleTaskRequest) model.TaskResponse
	Delete(model.HandleTaskRequest) model.TaskResponse
}

type TaskService interface {
	List(userId int) model.TaskResponse
	Create(userId int, payload model.TaskPayload) model.TaskResponse
	Update(id int, userId int, payload model.TaskPayload) model.TaskResponse
	Delete(id int, userId int) model.TaskResponse
}
