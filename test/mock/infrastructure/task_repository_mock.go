package infrastructure

import (
	"fmt"

	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

type taskRepositoryMock struct {
	shouldFail bool
	notFound   bool
}

func NewTaskRepositoryMock(shouldFail, notFound bool) contract.TaskRepository {
	return &taskRepositoryMock{shouldFail, notFound}
}

func (t taskRepositoryMock) FindAll() ([]model.Task, error) {
	var task []model.Task

	if t.shouldFail {
		return task, fmt.Errorf("internal server error")
	}

	if t.notFound {
		return task, nil
	}

	task = append(task, model.Task{
		Name:    "Hello World",
		Summary: "4ea5815b0d1d0d28f446d123f8e751cbca74c5c63a3cde51375c0fad8c947ff5fabc4885656cba0399d86155208d375844c7d3811bf544f99bc5335a",
		UserId:  1,
	})

	return task, nil
}

func (t taskRepositoryMock) FindByUserId(id int) ([]model.Task, error) {
	var task []model.Task

	if t.shouldFail {
		return task, fmt.Errorf("internal server error")
	}

	if t.notFound {
		return task, nil
	}

	task = append(task, model.Task{
		Name:    "Hello World",
		Summary: "4ea5815b0d1d0d28f446d123f8e751cbca74c5c63a3cde51375c0fad8c947ff5fabc4885656cba0399d86155208d375844c7d3811bf544f99bc5335a",
		UserId:  1,
	})

	return task, nil
}

func (t taskRepositoryMock) Create(task model.Task) (model.Task, error) {
	if t.shouldFail {
		return task, fmt.Errorf("internal server error")
	}

	return model.Task{}, nil
}

func (t taskRepositoryMock) Update(id int, userId int, task model.Task) (model.Task, error) {
	if t.shouldFail {
		return task, fmt.Errorf("internal server error")
	}

	return model.Task{}, nil
}

func (t taskRepositoryMock) Delete(id int) error {
	if t.shouldFail {
		return fmt.Errorf("internal server error")
	}

	return nil
}
