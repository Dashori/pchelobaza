package repository

import "backend/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetUserByLogin(login string) (*models.User, error)
	GetUserById(id uint64) (*models.User, error)
	UpdateUser(user *models.User) error
}
