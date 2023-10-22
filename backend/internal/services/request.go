package services

import "backend/internal/models"

type RequestService interface {
	GetAllRequests() ([]models.Request, error)
	GetRequestsPagination(limit int, skipped int) ([]models.Request, error)
	GetUserRequest(UserLogin string) (*models.Request, error)
	PatchUserRequest(request models.Request) error
}
