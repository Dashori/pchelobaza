package services

import "backend/internal/models"

type UserService interface {
	Create(newUser *models.User) (*models.User, error)
	Login(login, password string) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
	Update(user *models.UserPatch) (error)
	// SetRole() error
	// Create(client *models.Client, password string) (*models.Client, error)
	// Login(login, password string) (*models.Client, error)
	// GetClientById(id uint64) (*models.Client, error)
	// GetClientByLogin(login string) (*models.Client, error)
}