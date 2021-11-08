package repository

import (
	"log"

	"github.com/ViniciusMartinsS/manager/internal/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) domain.TaskRepository {
	return taskRepository{conn}
}

func (t taskRepository) FindAll() ([]domain.Task, error) {
	var task []domain.Task

	result := t.conn.
		Find(&task)

	if result.Error != nil {
		log.Println("[ERROR] Executing FindAll on Task repository: ", result.Error)
		return make([]domain.Task, 0), result.Error
	}

	return task, nil
}

func (t taskRepository) FindById(id int) (domain.Task, error) {
	var task domain.Task

	result := t.conn.Find(&task, id)

	if result.Error != nil {
		log.Println("[ERROR] Executing FindByUserId on Task repository: ", result.Error)
		return domain.Task{}, result.Error
	}

	return task, nil
}

func (t taskRepository) FindByUserId(id int) ([]domain.Task, error) {
	var task []domain.Task

	result := t.conn.
		Where("user_id = ?", id).
		Find(&task)

	if result.Error != nil {
		log.Println("[ERROR] Executing FindByUserId on Task repository: ", result.Error)
		return make([]domain.Task, 0), result.Error
	}

	return task, nil
}

func (t taskRepository) Create(task domain.Task) (domain.Task, error) {
	result := t.conn.
		Create(&task)

	if result.Error != nil {
		log.Println("[ERROR] Executing Create on Task repository: ", result.Error)
		return domain.Task{}, result.Error
	}

	return task, nil
}

func (t taskRepository) Update(id int, userId int, task domain.Task) (domain.Task, error) {
	result := t.conn.
		Model(&domain.Task{}).
		Where("id = ? AND user_id = ?", id, userId).
		Updates(task)

	if result.Error != nil {
		log.Println("[ERROR] Executing Update on Task repository: ", result.Error)
		return domain.Task{}, result.Error
	}

	return task, nil
}

func (t taskRepository) Delete(id int) (bool, error) {
	result := t.conn.
		Delete(&domain.Task{}, id)

	if result.Error != nil {
		log.Println("[ERROR] Executing Delete on Task repository: ", result.Error)
		return false, result.Error
	}

	return true, nil
}
