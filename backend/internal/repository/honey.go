package repository

import "backend/internal/models"

type HoneyRepository interface {
	GetAllHoney() ([]models.Honey, error)
}
