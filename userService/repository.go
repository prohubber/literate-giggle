package userService

import (
	"gorm.io/gorm"
)

// Repository provides methods to interact with the Users table
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates a new user repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// GetAll retrieves all users
func (r *Repository) GetAll() ([]User, error) {
	var users []User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Create adds a new user to the database
func (r *Repository) Create(user *User) error {
	return r.DB.Create(user).Error
}

// Update modifies an existing user
func (r *Repository) Update(user *User) error {
	return r.DB.Save(user).Error
}

// Delete removes a user by ID
func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&User{}, id).Error
}
