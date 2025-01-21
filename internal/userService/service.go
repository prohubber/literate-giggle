package userService

import "main/project/internal/taskService"

// UserService provides business logic for managing users
type UserService struct {
	repo *Repository
}

func (s *UserService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	if err := s.repo.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// NewUserService creates a new user service
func NewUserService(repo *Repository) *UserService {
	return &UserService{repo: repo}
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

// CreateUser adds a new user
func (s *UserService) CreateUser(user *User) error {
	return s.repo.Create(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *User) error {
	return s.repo.Update(user)
}

// DeleteUserByID deletes a user by ID
func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.Delete(id)
}
