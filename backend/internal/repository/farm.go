package repository

import "backend/internal/models"

type FarmRepository interface {
	CreateFarm(newFarm *models.Farm) error
	GetFarmByName(name string) (*models.Farm, error)
	GetFarmById(id uint64) (*models.Farm, error)
	GetUsersFarm(userId uint64, limit int, skipped int) ([]models.Farm, error)
	PatchFarm(newFarm *models.Farm) error
}
