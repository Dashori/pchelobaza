package repository

import "backend/internal/models"

type HoneyRepository interface {
	GetAllHoney() ([]models.Honey, error)
	GetHoneyId(name string) (uint64, error)
	GetFarmHoney(name string) ([]models.Honey, error)
}
