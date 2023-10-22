package repository

import "backend/internal/models"

type RequestRepository interface {
	Create(request *models.Request) error
	GetAllRequests() ([]models.Request, error)
	GetRequestsPagination(limit int, skipped int) ([]models.Request, error)
	GetUserRequest(UserLogin string) (*models.Request, error)
	PatchUserRequest(request *models.Request) error
}
