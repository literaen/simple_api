package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)

	// GetAllTasks
	// возвращаем массив из всех задач и ошибку
	GetAllTasks() ([]Task, error)

	// // GetTasksByUserID - Передаем ID
	// // возвращаем массив из всех задач юзера и ошибку
	// GetTasksByUserID(id uint) ([]Task, error)

	// UpdateTaskByID - Передаем id и Task
	// возвращаем обновленный Task и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)

	// DeleteTaskByID - Передаем id для удаления
	// возвращаем только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// func (r *taskRepository) GetTasksByUserID(id uint) ([]Task, error) {
// 	var tasks []Task
// 	err := r.db.Where("user_id = ?", id).Find(&tasks).Error
// 	return tasks, err
// }

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	result := r.db.Model(&Task{}).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}

	var updatedTask Task
	if err := r.db.First(&updatedTask, id).Error; err != nil {
		return Task{}, err
	}

	return updatedTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task Task
	result := r.db.Where("id = ?", id).Delete(&task)
	return result.Error
}
