package repository

import "backend/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	// GetUserByLogin(login string) (*models.User, error)
	// GetClientById(id uint64) (*models.Client, error)
	// GetAllClient() ([]models.Client, error)
	// Delete(id uint64) error
}