package services

import "backend/internal/models"

type ConferenceService interface {
	GetAllConferences(limit int, skipped int) ([]models.Conference, error)
	CreateConference(conference *models.Conference) (*models.Conference, error)
	GetConferenceByName(name string) (*models.Conference, error)
	PatchConference(conference *models.Conference) error
	GetAllConferenceUsers(name string) ([]models.User, error)
	GetConferenceUsers(name string, limit int, skipped int) ([]models.User, error)
	PatchConferenceUsers(name string, login string) error
	GetConferenceReviews(name string, limit int, skipped int) ([]models.Review, error)
	CreateReview(review *models.Review) (*models.Review, error)
}
