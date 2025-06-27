package services

import "online-school/internal/models"

// UserService defines methods related to user operations.
type UserService interface {
	CreateUser(user models.User) (models.User, error)
	Authenticate(email, password string) (string, error)
}

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	// Add other methods as needed, e.g., Create, Update, Delete
}

type userService struct {
	repo UserRepository
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}
