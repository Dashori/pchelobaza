package repository

import "backend/internal/models"

type ConferenceRepository interface {
	GetAllConferences(limit int, skipped int) ([]models.Conference, error)
	CreateConference(conference *models.Conference) error
	GetConferenceByName(name string) (*models.Conference, error)
	PatchConference(conference *models.Conference) error
	GetAllConferenceUsers(name string) ([]models.User, error)
	GetConferenceUsers(name string, limit int, skipped int) ([]models.User, error)
	PatchConferenceUsers(conference *models.Conference, UserId uint64) error
	GetConferenceReviews(name string, limit int, skipped int) ([]models.Review, error)
	CreateReview(review *models.Review) error
}
