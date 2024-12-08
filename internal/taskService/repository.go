package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)

	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)

	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)

	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error

	// PatchTaskByID - Частичное обновление задачи по ID
	PatchTaskByID(id uint, task Task) (Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// CreateTask - Создает новую задачу в БД
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

// GetAllTasks - Возвращает все задачи из БД
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// UpdateTaskByID - Обновляет задачу по ID
func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task
	// Находим задачу по ID
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return Task{}, err
	}
	// Обновляем поля задачи
	task.ID = existingTask.ID
	result := r.db.Save(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

// DeleteTaskByID - Удаляет задачу по ID
func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task Task
	// Находим задачу по ID
	if err := r.db.First(&task, id).Error; err != nil {
		return err
	}
	// Удаляем задачу
	result := r.db.Delete(&task)
	return result.Error
}

// PatchTaskByID - Частичное обновление задачи по ID
func (r *taskRepository) PatchTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task
	// Находим задачу по ID
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return Task{}, err
	}

	// Обновляем только измененные поля
	result := r.db.Model(&existingTask).Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}

	// Возвращаем обновленную задачу
	return existingTask, nil
}
