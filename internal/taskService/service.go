package taskService

// TaskService — структура для сервиса
type TaskService struct {
	repo TaskRepository
}

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

// NewService — конструктор для создания нового сервиса
func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask — вызывает метод репозитория для создания задачи
func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

// GetAllTasks — вызывает метод репозитория для получения всех задач
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

// UpdateTaskByID — вызывает метод репозитория для обновления задачи по ID
func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

// PatchTaskByID — вызывает метод репозитория для частичного обновления задачи по ID
func (s *TaskService) PatchTaskByID(id uint, task Task) (Task, error) {
	return s.repo.PatchTaskByID(id, task)
}

// DeleteTaskByID — вызывает метод репозитория для удаления задачи по ID
func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
