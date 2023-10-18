package repository

import "backend/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetUserByLogin(login string) (*models.User, error)
	UpdateUser(user *models.UserPatch) error
}
