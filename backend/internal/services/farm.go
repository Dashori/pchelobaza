package services

import "backend/internal/models"

type FarmService interface {
	CreateFarm(newFarm *models.Farm) (*models.Farm, error)
	GetFarm(name string) (*models.Farm, error)
	GetUsersFarm(login string, limit int, skipped int) ([]models.Farm, error)
	PatchFarm(newFarm *models.Farm) error
}
