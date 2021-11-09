package repository

import (
	"fmt"
	"log"

	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"gorm.io/gorm"
)

type taskRepository struct {
	conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) contract.TaskRepository {
	return taskRepository{conn}
}

func (t taskRepository) FindAll() ([]model.Task, error) {
	var task []model.Task

	result := t.conn.
		Find(&task)

	if result.Error != nil {
		log.Println("[ERROR] Executing FindAll on Task repository: ", result.Error)
		return make([]model.Task, 0), result.Error
	}

	return task, nil
}

func (t taskRepository) FindById(id int) (model.Task, error) {
	var task model.Task

	result := t.conn.Find(&task, id)

	if result.Error != nil {
		log.Println("[ERROR] Executing FindByUserId on Task repository: ", result.Error)
		return model.Task{}, result.Error
	}

	return task, nil
}

func (t taskRepository) FindByUserId(id int) ([]model.Task, error) {
	var task []model.Task

	result := t.conn.
		Where("user_id = ?", id).
		Find(&task)

	if result.Error != nil {
		log.Println("[ERROR] Executing FindByUserId on Task repository: ", result.Error)
		return make([]model.Task, 0), result.Error
	}

	return task, nil
}

func (t taskRepository) Create(task model.Task) (model.Task, error) {
	result := t.conn.
		Create(&task)

	if result.Error != nil {
		log.Println("[ERROR] Executing Create on Task repository: ", result.Error)
		return model.Task{}, result.Error
	}

	return task, nil
}

func (t taskRepository) Update(id int, userId int, task model.Task) (model.Task, error) {
	result := t.conn.
		Model(&model.Task{}).
		Where("id = ? AND user_id = ?", id, userId).
		Updates(task)

	if result.Error != nil {
		log.Println("[ERROR] Executing Update on Task repository: ", result.Error)
		return model.Task{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Task{}, fmt.Errorf("record not found")
	}

	return task, nil
}

func (t taskRepository) Delete(id int) error {
	result := t.conn.
		Delete(&model.Task{}, id)

	if result.Error != nil {
		log.Println("[ERROR] Executing Delete on Task repository: ", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}

	return nil
}
