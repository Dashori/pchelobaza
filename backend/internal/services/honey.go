package services

import "backend/internal/models"

type HoneyService interface {
	GetAllHoney() ([]models.Honey, error)
}
